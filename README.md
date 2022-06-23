# b64headerdecoder
Base64 Header Decoder

Usage:  ./bin/b64hd  [-h, --help] [inputfile] [outputfile]
Reads input stream, finds 'X-ImunifyEmail-Filter-Info:' following base64 encoded block. 
Decode the block to output stream and print other data without changes