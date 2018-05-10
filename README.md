# Amazon Kinesis Stream consumer and producer example using aws-sdk-go

Demonstration for

* create stream
* wait until Kinesis stream available
* describe stream
* put record
* put multiple records (using PutRecords API)
* get records using shard iterator
* delete stream

Authorization is depends on aws-sdk-go. I use `~/.aws/config` for this purpose.

Example output

```
$ go run main.go
2018/05/10 09:39:43 main.go:26: region eu-central-1
2018/05/10 09:39:43 main.go:58: CreateStream your-stream
2018/05/10 09:39:43 main.go:68: CreateStream output: {

}
2018/05/10 09:39:43 main.go:69: WaitUntilStreamExists
2018/05/10 09:40:04 main.go:75: describing stream {
  StreamName: "your-stream"
}
2018/05/10 09:40:04 main.go:81: {
  StreamDescription: {
    EncryptionType: "NONE",
    EnhancedMonitoring: [{
        ShardLevelMetrics: []
      }],
    HasMoreShards: false,
    RetentionPeriodHours: 24,
    Shards: [{
        HashKeyRange: {
          EndingHashKey: "340282366920938463463374607431768211455",
          StartingHashKey: "0"
        },
        SequenceNumberRange: {
          StartingSequenceNumber: "49584328888477368320818992204973547556223369515045486594"
        },
        ShardId: "shardId-000000000000"
      }],
    StreamARN: "arn:aws:kinesis:eu-central-1:xxxxxxxxx:stream/your-stream",
    StreamCreationTimestamp: 2018-05-10 06:39:43 +0000 UTC,
    StreamName: "your-stream",
    StreamStatus: "ACTIVE"
  }
}
2018/05/10 09:40:04 main.go:87: putting record {
  Data: <binary> len 4,
  PartitionKey: "key1",
  StreamName: "your-stream"
}
2018/05/10 09:40:04 main.go:93: putOutput: {
  SequenceNumber: "49584328888477368320818992204989263591878361274864631810",
  ShardId: "shardId-000000000000"
}
2018/05/10 09:40:04 main.go:109: records to put: {
  Records: [
    {
      Data: <binary> len 5,
      PartitionKey: "key2"
    },
    {
      Data: <binary> len 5,
      PartitionKey: "key2"
    },
    {
      Data: <binary> len 5,
      PartitionKey: "key2"
    },
    {
      Data: <binary> len 5,
      PartitionKey: "key2"
    },
    {
      Data: <binary> len 5,
      PartitionKey: "key2"
    },
    {
      Data: <binary> len 5,
      PartitionKey: "key2"
    },
    {
      Data: <binary> len 5,
      PartitionKey: "key2"
    },
    {
      Data: <binary> len 5,
      PartitionKey: "key2"
    },
    {
      Data: <binary> len 5,
      PartitionKey: "key2"
    },
    {
      Data: <binary> len 5,
      PartitionKey: "key2"
    }
  ],
  StreamName: "your-stream"
}
2018/05/10 09:40:04 main.go:116: putsOutput: {
  FailedRecordCount: 0,
  Records: [
    {
      SequenceNumber: "49584328888477368320818992204990472517697975972758814722",
      ShardId: "shardId-000000000000"
    },
    {
      SequenceNumber: "49584328888477368320818992204991681443517590601933520898",
      ShardId: "shardId-000000000000"
    },
    {
      SequenceNumber: "49584328888477368320818992204992890369337205231108227074",
      ShardId: "shardId-000000000000"
    },
    {
      SequenceNumber: "49584328888477368320818992204994099295156819860282933250",
      ShardId: "shardId-000000000000"
    },
    {
      SequenceNumber: "49584328888477368320818992204995308220976434489457639426",
      ShardId: "shardId-000000000000"
    },
    {
      SequenceNumber: "49584328888477368320818992204996517146796049118632345602",
      ShardId: "shardId-000000000000"
    },
    {
      SequenceNumber: "49584328888477368320818992204997726072615663747807051778",
      ShardId: "shardId-000000000000"
    },
    {
      SequenceNumber: "49584328888477368320818992204998934998435278376981757954",
      ShardId: "shardId-000000000000"
    },
    {
      SequenceNumber: "49584328888477368320818992205000143924254893006156464130",
      ShardId: "shardId-000000000000"
    },
    {
      SequenceNumber: "49584328888477368320818992205001352850074507635331170306",
      ShardId: "shardId-000000000000"
    }
  ]
}
2018/05/10 09:40:04 main.go:126: GetShardIterator input {
  ShardId: "shardId-000000000000",
  ShardIteratorType: "TRIM_HORIZON",
  StreamName: "your-stream"
}
2018/05/10 09:40:04 main.go:133: iteratorOutput: {
  ShardIterator: "AAAAAAAAAAHV462OY+YY5LwnaEf1owMikdwBgO8qf9M6MKPqC96PsqMGdVcqA1T5doX/hcBHw0Gw6pzCmkC23lBuyjaevopyL9px5xJ5D9j/9oHojRnD7E2PJ0FgLoelbtBqoU9qly8htHMiBbjBOilHqRmJKpd/YLrkm/rRAXGWgMdmH34LEKsLp11eGYmtRBlUtmoj1Wy97e4Folla1aT5WtgegLz/"
}
2018/05/10 09:40:04 main.go:134: GetRecords by ShardIterator AAAAAAAAAAHV462OY+YY5LwnaEf1owMikdwBgO8qf9M6MKPqC96PsqMGdVcqA1T5doX/hcBHw0Gw6pzCmkC23lBuyjaevopyL9px5xJ5D9j/9oHojRnD7E2PJ0FgLoelbtBqoU9qly8htHMiBbjBOilHqRmJKpd/YLrkm/rRAXGWgMdmH34LEKsLp11eGYmtRBlUtmoj1Wy97e4Folla1aT5WtgegLz/
2018/05/10 09:40:04 main.go:143: records: {
  MillisBehindLatest: 0,
  NextShardIterator: "AAAAAAAAAAFSUvLdPWDxYtGUc6W94oReH0/kxqzMB/KJh213VVg/aJv5dxunh1cGK6/kY5ZVN1fhLK63p51sAOmX3F28Ax7yC4MTDdHIgAUCF3M4dg4wypVI5i5CA6kpNxHsmHvo9SB7AQf5obpQOBPXdahAK7iB/Tk6K4ESXD9s56yGxPmYAZU8g4sbgENHn+I3I5ZS2M1/yH9yurgeMG/tY0/+ORZ1",
  Records: [
    {
      ApproximateArrivalTimestamp: 2018-05-10 06:40:04 +0000 UTC,
      Data: <binary> len 4,
      PartitionKey: "key1",
      SequenceNumber: "49584328888477368320818992204989263591878361274864631810"
    },
    {
      ApproximateArrivalTimestamp: 2018-05-10 06:40:04 +0000 UTC,
      Data: <binary> len 5,
      PartitionKey: "key2",
      SequenceNumber: "49584328888477368320818992204990472517697975972758814722"
    },
    {
      ApproximateArrivalTimestamp: 2018-05-10 06:40:04 +0000 UTC,
      Data: <binary> len 5,
      PartitionKey: "key2",
      SequenceNumber: "49584328888477368320818992204991681443517590601933520898"
    },
    {
      ApproximateArrivalTimestamp: 2018-05-10 06:40:04 +0000 UTC,
      Data: <binary> len 5,
      PartitionKey: "key2",
      SequenceNumber: "49584328888477368320818992204992890369337205231108227074"
    },
    {
      ApproximateArrivalTimestamp: 2018-05-10 06:40:04 +0000 UTC,
      Data: <binary> len 5,
      PartitionKey: "key2",
      SequenceNumber: "49584328888477368320818992204994099295156819860282933250"
    },
    {
      ApproximateArrivalTimestamp: 2018-05-10 06:40:04 +0000 UTC,
      Data: <binary> len 5,
      PartitionKey: "key2",
      SequenceNumber: "49584328888477368320818992204995308220976434489457639426"
    },
    {
      ApproximateArrivalTimestamp: 2018-05-10 06:40:04 +0000 UTC,
      Data: <binary> len 5,
      PartitionKey: "key2",
      SequenceNumber: "49584328888477368320818992204996517146796049118632345602"
    },
    {
      ApproximateArrivalTimestamp: 2018-05-10 06:40:04 +0000 UTC,
      Data: <binary> len 5,
      PartitionKey: "key2",
      SequenceNumber: "49584328888477368320818992204997726072615663747807051778"
    },
    {
      ApproximateArrivalTimestamp: 2018-05-10 06:40:04 +0000 UTC,
      Data: <binary> len 5,
      PartitionKey: "key2",
      SequenceNumber: "49584328888477368320818992204998934998435278376981757954"
    },
    {
      ApproximateArrivalTimestamp: 2018-05-10 06:40:04 +0000 UTC,
      Data: <binary> len 5,
      PartitionKey: "key2",
      SequenceNumber: "49584328888477368320818992205000143924254893006156464130"
    },
    {
      ApproximateArrivalTimestamp: 2018-05-10 06:40:04 +0000 UTC,
      Data: <binary> len 5,
      PartitionKey: "key2",
      SequenceNumber: "49584328888477368320818992205001352850074507635331170306"
    }
  ]
}
2018/05/10 09:40:04 main.go:146: getting next records using records.NextShardIterator AAAAAAAAAAFSUvLdPWDxYtGUc6W94oReH0/kxqzMB/KJh213VVg/aJv5dxunh1cGK6/kY5ZVN1fhLK63p51sAOmX3F28Ax7yC4MTDdHIgAUCF3M4dg4wypVI5i5CA6kpNxHsmHvo9SB7AQf5obpQOBPXdahAK7iB/Tk6K4ESXD9s56yGxPmYAZU8g4sbgENHn+I3I5ZS2M1/yH9yurgeMG/tY0/+ORZ1
2018/05/10 09:40:05 main.go:154: recordsSecond {
  MillisBehindLatest: 0,
  NextShardIterator: "AAAAAAAAAAHMMGVnDg6G2/vJVXMOfxb3G5fS5V8uJNxuzyZkwhtnkxANt0t9fDXSSbePJzqCn+WzXuEkNObkFdZXTVlRmKrFzvWJk8Qmm1N/8xR0KDJuemyQlWWsrc75U0w0TzgT9PrFClqjzXsyqokx3fj6Cxs66tXf5kIcIhCtYxS2FRVyVQHISGbMxBF326ZRYvJS0XyKuak5z8w8imSpNBq8GSvZ",
  Records: []
}
2018/05/10 09:40:05 main.go:42: deleting the stream {
  StreamName: "your-stream"
}
2018/05/10 09:40:05 main.go:49: deleteOutput: {

}
```

For handling errors, see [handling errors · aws/aws-sdk-go Wiki](https://github.com/aws/aws-sdk-go/wiki/handling-errors).

## LICENSE

MIT

forked from https://github.com/ypapax/amazon-kinesis-go-example
