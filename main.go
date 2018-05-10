package main

import (
	"flag"
	"fmt"
	"log"

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

	s := session.New(&aws.Config{Region: aws.String(*region)})
	kc := kinesis.New(s)

	streamName := aws.String(*stream)

	log.Println("streamName", *streamName)

	out, err := kc.CreateStream(&kinesis.CreateStreamInput{
		ShardCount: aws.Int64(1),
		StreamName: streamName,
	})
	if err != nil {
		panic(err)
	}
	log.Printf("CreateStream output: %v\n", out)
	log.Println("WaitUntilStreamExists")
	if err := kc.WaitUntilStreamExists(&kinesis.DescribeStreamInput{StreamName: streamName}); err != nil {
		panic(err)
	}
	describeStreamInput := kinesis.DescribeStreamInput{StreamName: streamName}
	log.Println("describing stream %+v", describeStreamInput)
	streams, err := kc.DescribeStream(&describeStreamInput)
	if err != nil {
		panic(err)
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
		panic(err)
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
		panic(err)
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
	log.Println("iteratorInput %+v", iteratorInput)
	// retrieve iterator
	iteratorOutput, err := kc.GetShardIterator(&iteratorInput)
	if err != nil {
		panic(err)
	}
	log.Printf("iteratorOutput: %v\n", iteratorOutput)
	log.Printf("getting records by iterator %+v", iteratorOutput.ShardIterator)
	// get records use shard iterator for making request
	records, err := kc.GetRecords(&kinesis.GetRecordsInput{
		ShardIterator: iteratorOutput.ShardIterator,
	})
	if err != nil {
		panic(err)
	}
	log.Printf("records: %v\n", records)

	// and, you can iteratively make GetRecords request using records.NextShardIterator
	log.Printf("getting next records using records.NextShardIterator %+v", records.NextShardIterator)
	recordsSecond, err := kc.GetRecords(&kinesis.GetRecordsInput{
		ShardIterator: records.NextShardIterator,
	})
	if err != nil {
		panic(err)
	}
	log.Printf("recordsSecond %v\n", recordsSecond)
	deleteParams := kinesis.DeleteStreamInput{
		StreamName: streamName,
	}
	log.Printf("deleting the stream %+v\n", deleteParams)
	// OK, finally delete your stream
	deleteOutput, err := kc.DeleteStream(&deleteParams)
	if err != nil {
		panic(err)
	}
	log.Printf("deleteOutput: %v\n", deleteOutput)
}
