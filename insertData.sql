CREATE  TABLE `library`.`books` (
  `book_ID` VARCHAR(255) NOT NULL,
  `book_img` varchar(255),
  `book_title` VARCHAR(45) NOT NULL,
  `book_category` VARCHAR(45) NOT NULL,
  `book_author` VARCHAR(45) NOT NULL,
  `book_publisher` VARCHAR(45) NOT NULL,
  `books_year` INT NOT NULL,
  `book_status` INT NOT NULL,
  PRIMARY KEY (`book_ID`),
  UNIQUE INDEX `book_ID_UNIQUE` (`book_ID` ASC) VISIBLE);


INSERT INTO books (book_ID,book_img, book_title, book_category, book_author, book_publisher, books_year, book_status)
VALUES
  ("book-00","Library-frontend/public/images/Mathematic.jpg",'คณิตศาสตร์', 'แบบฝึกหัดคณิตศาสตร์ ป.6', 'สถาบันส่งเสริมการสอนวิทยาศาสตร์และเทคโนโลยี', 'สสวท.', 2024, 0),
  ("book-01","Library-frontend/public/images/Science.jpg",'วิทยาศาสตร์', 'วิทยาศาสตร์รอบตัว', 'กลุ่มนักเขียน', 'สำนักพิมพ์อมรินทร์', 2023, 0),
  ("book-02","Library-frontend/public/images/Thai.jpg",'ภาษาไทย', 'ไวยากรณ์ภาษาไทย ม.3', 'อาจารย์สุภัทร สุวรรณ', 'พิงค์กีส์การศึกษา', 2022, 1)
  ;
  