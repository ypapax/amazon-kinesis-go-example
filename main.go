package main

import (
	"flag"
	"fmt"
	"log"

	"os"
	"os/signal"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

var (
	stream = flag.String("stream", "your-stream", "your stream name")
	region = flag.String("region", "eu-central-1", "your AWS region")
)

func main() {
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	awsConfig := &aws.Config{Region: aws.String(*region)}
	log.Println("region", *awsConfig.Region)
	s, err := session.NewSession(awsConfig)
	if err != nil {
		log.Println(err)
		return
	}
	kc := kinesis.New(s)

	streamName := aws.String(*stream)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	delete := func() {
		deleteParams := kinesis.DeleteStreamInput{
			StreamName: streamName,
		}
		log.Printf("deleting the stream %+v\n", deleteParams)
		// OK, finally delete your stream
		deleteOutput, err := kc.DeleteStream(&deleteParams)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("deleteOutput: %v\n", deleteOutput)
	}
	go func() {
		for sig := range c {
			log.Println("exiting", sig)
			delete()
			os.Exit(1)
		}
	}()
	log.Println("CreateStream", *streamName)
	out, err := kc.CreateStream(&kinesis.CreateStreamInput{
		ShardCount: aws.Int64(1),
		StreamName: streamName,
	})
	if err != nil {
		log.Println(err)
		return
	}
	defer delete()
	log.Printf("CreateStream output: %v\n", out)
	log.Println("WaitUntilStreamExists")
	if err := kc.WaitUntilStreamExists(&kinesis.DescribeStreamInput{StreamName: streamName}); err != nil {
		log.Println(err)
		return
	}
	describeStreamInput := kinesis.DescribeStreamInput{StreamName: streamName}
	log.Println("describing stream", describeStreamInput)
	streams, err := kc.DescribeStream(&describeStreamInput)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("%v\n", streams)
	record := kinesis.PutRecordInput{
		Data:         []byte("hoge"),
		StreamName:   streamName,
		PartitionKey: aws.String("key1"),
	}
	log.Printf("putting record %+v\n", record)
	putOutput, err := kc.PutRecord(&record)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("putOutput: %v\n", putOutput)

	// put 10 records using PutRecords API
	entries := make([]*kinesis.PutRecordsRequestEntry, 10)
	for i := 0; i < len(entries); i++ {
		entries[i] = &kinesis.PutRecordsRequestEntry{
			Data:         []byte(fmt.Sprintf("hoge%d", i)),
			PartitionKey: aws.String("key2"),
		}
	}

	putRecordsInput := kinesis.PutRecordsInput{
		Records:    entries,
		StreamName: streamName,
	}

	log.Printf("records to put: %v\n", putRecordsInput)
	putsOutput, err := kc.PutRecords(&putRecordsInput)
	if err != nil {
		log.Println(err)
		return
	}
	// putsOutput has Records, and its shard id and sequece enumber.
	log.Printf("putsOutput: %v\n", putsOutput)

	iteratorInput := kinesis.GetShardIteratorInput{
		// Shard Id is provided when making put record(s) request.
		ShardId:           putOutput.ShardId,
		ShardIteratorType: aws.String("TRIM_HORIZON"),
		// ShardIteratorType: aws.String("AT_SEQUENCE_NUMBER"),
		// ShardIteratorType: aws.String("LATEST"),
		StreamName: streamName,
	}
	log.Println("GetShardIterator input", iteratorInput)
	// retrieve iterator
	iteratorOutput, err := kc.GetShardIterator(&iteratorInput)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("iteratorOutput: %v\n", iteratorOutput)
	log.Printf("GetRecords by ShardIterator %+v", *iteratorOutput.ShardIterator)
	// get records use shard iterator for making request
	records, err := kc.GetRecords(&kinesis.GetRecordsInput{
		ShardIterator: iteratorOutput.ShardIterator,
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("records: %v\n", records)

	// and, you can iteratively make GetRecords request using records.NextShardIterator
	log.Printf("getting next records using records.NextShardIterator %+v", *records.NextShardIterator)
	recordsSecond, err := kc.GetRecords(&kinesis.GetRecordsInput{
		ShardIterator: records.NextShardIterator,
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("recordsSecond %v\n", recordsSecond)

}