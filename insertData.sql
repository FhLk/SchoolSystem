CREATE TABLE `library`.`books_table` (
  `book_ID` VARCHAR(255) NOT NULL,
  `book_title` VARCHAR(45) NOT NULL,
  `book_category` VARCHAR(45) NOT NULL,
  `book_author` VARCHAR(45) NOT NULL,
  `book_publisher` VARCHAR(45) NOT NULL,
  `books_year` INT NOT NULL,
  `book_status` TINYINT NOT NULL,
  PRIMARY KEY (`book_ID`),
  UNIQUE INDEX `book_ID_UNIQUE` (`book_ID` ASC) VISIBLE);