# trnsl8
trnsl8 is a CLI written in Go (using Cobra package) that allows users to translate text and detect the language of input through the use of AWS Translate and AWS Comprehend

## Usage

```sh
# Translate text into a target language
# More information on language codes: https://docs.aws.amazon.com/translate/latest/dg/what-is.html
trnsl8 to <language-code> <text-to-translate>  

# Detect what language a text is in
trnsl8 detect <text-to-detect> 

```


## Examples

```sh
# Translation
trnsl8 to jp "I want this sentence in japanese"
trnsl8 to en "Ich möchte diesen Satz auf Englisch"

# Language detection
trnsl8 detect "ce limbă este aceasta?" # hint: its romanian!
```

## Next steps
* Properly handle JSON returned from AWS 
* Implement -f flag to allow user to input files
* ...
