# Base64 Header Decoder

Usage:

```b64hd  [-h, --help] [inputfile] [outputfile]```

Reads input stream, finds 'X-ImunifyEmail-Filter-Info:' following base64 encoded block. 
Decodes the block to output stream and prints other data without changes
