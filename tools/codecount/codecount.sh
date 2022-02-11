#!/bin/bash
cd ../../../
output_file="sgn-v2/tools/codecount/code-count-report.txt"
echo "Deleting existing report..."
rm $output_file
touch $output_file
echo "Generating report..."
echo "SGN-V2 CODE LINE COUNT REPORT" >> $output_file
dt=$(date '+%m/%d/%Y %H:%M:%S');
echo "$dt" >> $output_file

echo "" >> $output_file
echo "The code count in this report only includes folders and files with effective code, excluding all folders and files which are either documentations, automatically generated, or for testing purposes." >> $output_file

INCLUDED_FOLDERS="sgn-v2/app sgn-v2/common sgn-v2/eth sgn-v2/executor sgn-v2/proto sgn-v2/relayer sgn-v2/x/cbridge sgn-v2/x/distribution sgn-v2/x/farming sgn-v2/x/gov sgn-v2/x/message sgn-v2/x/mint sgn-v2/x/pegbridge sgn-v2/x/slashing sgn-v2/x/staking sgn-v2/x/sync"
EXCLUDED_FILENAME_REGEX='pb.go|.txt|test|bindings|.md'

str_to_replace=" "
replace_str='|'
INCLUDED_FOLDERS_REGEX=${INCLUDED_FOLDERS//$str_to_replace/$replace_str}

replace_str="\n"
INCLUDED_FOLDERS_TO_PRINT=${INCLUDED_FOLDERS//$str_to_replace/$replace_str}

echo "" >> $output_file
echo "Here is the total effective code count for the whole project." >> $output_file
echo "The included folders are:" >> $output_file
printf $INCLUDED_FOLDERS_TO_PRINT >> $output_file
echo "" >> $output_file
gocloc --not-match=$EXCLUDED_FILENAME_REGEX  --match-d=$INCLUDED_FOLDERS_REGEX . >> $output_file
echo "" >> $output_file

echo "" >> $output_file
echo "Below is the breakdown of the code count for each folder mentioned above:" >> $output_file
echo "" >> $output_file

for FOLDER in $INCLUDED_FOLDERS
do
echo "" >> $output_file
echo $FOLDER >> $output_file
gocloc --not-match=$EXCLUDED_FILENAME_REGEX  --match-d=$FOLDER . >> $output_file
echo "" >> $output_file
done

echo "Done. Report generated."