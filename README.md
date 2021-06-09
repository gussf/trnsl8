# trnsl8
trnsl8 is a CLI written in Go that allows users to translate text and detect the language of input through the use of AWS Translate and AWS Comprehend


## Examples

```sh
# Translation
trnsl8 to jp "I want this sentence in japanese"
trnsl8 to en "Ich möchte diesen Satz auf Englisch"

# Language detection
trnsl8 detect "ce limbă este aceasta?" # hint: its romanian!
```

## Next steps
* Implement -f flag to allow user to input files
* ...
