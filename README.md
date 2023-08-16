## Text-WaterMark

Add watermark to your text,writing in go


### How it works?

[Be careful what you copy: Invisibly inserting usernames into text with Zero-Width Characters](https://medium.com/@umpox/be-careful-what-you-copy-invisibly-inserting-usernames-into-text-with-zero-width-characters-18b4e6f17b66)

### Usage

```
go get -u github.com/MoYoez/text-watermark
```

### Example

```Go
    rawName := "我能吞下玻璃而不傷身體"
	encName := "WaterMark"
	encText := AddWaterMarkToText(rawName, encName)
	fmt.Print(getWaterMark(encText))

```


### License

MIT
