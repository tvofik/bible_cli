New Testament
---
Genesis
Exodus
Leviticus
Numbers
Deuteronomy
Joshua
Judges
Ruth
1 Samuel
2 Samuel
1 Kings
2 Kings
1 Chronicles
2 Chronicles
Ezra
Nehemiah
Esther
Job
Psalm
Proverbs
Ecclesiastes
Song of Solomon
Isaiah
Jeremiah
Lamentations
Ezekiel
Daniel
Hosea
Joel
Amos
Obadiah
Jonah
Micah
Nahum
Habakkuk
Zephaniah
Haggai
Zechariah
Malachi

Old Testament
---
Matthew
Mark
Luke
John
Acts
Romans
1 Corinthians
2 Corinthians
Galatians
Ephesians
Philippians
Colossians
1 Thessalonians
2 Thessalonians
1 Timothy
2 Timothy
Titus
Philemon
Hebrews
James
1 Peter
2 Peter
1 John
2 John
3 John
Jude
Revelation


<!-- (^\s*[1-3]?\s*[A-Za-z]+)\s*(\d+)[;: ]+(\d+)([-, ]+\d+)? # Main Regex # -->
<!-- (^\s*[1-3]?\s*)([A-Za-z]+)\s*(\d+)[;: ]+(\d+)([-, ]+)?(\d+)? # Main Regex w/ Numbers before the books having a capture group & delimiters -->

<!-- (?P<book>^\s*[1-3]?\s*[A-Za-z]+)\s*(?P<chapter>\d+)[;: ]+(?P<verse>\d+)(?P<verses>[-, ]+\d+)? # Main Regex # -->
<!-- (?P<book>^\s*[1-3]?\s*[A-Za-z]+)\s*(?P<chapter>\d+)[;: ]+(?P<verse>\d+)(?P<delimiter>[-, ]+)?(?P<verses>\d+)?  #Main Regex w/ delimiter as a capture group -->

<!-- (?P<num>^\s*[1-3]?\s*)(?P<book>[A-Za-z]+)\s*(?P<chapter>\d+)[;: ]+(?P<verse>\d+)(?P<delimiter>[-, ]+)?(?P<extraverse>\d+)? # Main Regex w/ Numbers before the books having a capture group & delimiters -->









  Test Cases 
<!-- 
         1          john 3:10123343343545
 1 john 3:10123343343545
1 john           31442343232442212312         102392344231 10 11
john 3:10-12
JOHN
       John
john
John 3:1034322342355353
John3:1
john :2 # check
John 3 :       1
john 3:10,11,12
john 3 12 13

John 3:3-5
1 John 4:5 - 6
2 john 4:5 - 4:6
3 john 4:5 - 3 John 4:6
John 4:5 - 6
john 4:5 - 4:6
John 4:5 - 1 John 4:6
1john4:6
john 4
john 4-5
1 john 4-5
john3:10
john 3 10
john 3:10
john3:10-11
john 3 10-11
john 3:10-11
john 3:10,21
 -->


someone passes the passage 
- standardize it 
- check if it is valid & and contains everything 
- return the passage


Notes
- find a way to combine GetBookInfo and Contains 





John 3:10-11



{ 
  0
  0
  0
  []
}

<!-- ========================== -->
Flag: 
  --translation=nkjv | kjv | 

bible_cli open "John 10:10" -t kjv 


