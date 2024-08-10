CREATE TABLE IF NOT EXISTS `Library`.`Books` (
  `Id` VARCHAR(255) NOT NULL,
  `Img` VARCHAR(225) NOT NULL,
  `Title` VARCHAR(45) NOT NULL,
  `Category` VARCHAR(45) NOT NULL,
  `Author` VARCHAR(45) NOT NULL,
  `Publisher` VARCHAR(45) NOT NULL,
  `Year` INT NOT NULL,
  `Status` INT NOT NULL DEFAULT 0,
  PRIMARY KEY (`Id`))
ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `Library`.`Members` (
  `Id` VARCHAR(255) NOT NULL,
  `First_name` VARCHAR(255) NOT NULL,
  `Last_name` VARCHAR(255) NOT NULL,
  `Class` VARCHAR(45) NOT NULL,
  `Exp_card` DATE NOT NULL,
  PRIMARY KEY (`Id`))
ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `Library`.`Librarians` (
  `Id` VARCHAR(225) NOT NULL,
  `First_name` VARCHAR(45) NOT NULL,
  `Last_name` VARCHAR(45) NOT NULL,
  `Role` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`Id`))
ENGINE = InnoDB;

drop table books;
drop table members;
drop table librarians;

INSERT INTO books (Id,Img, Category,Title, Author, Publisher, Year, Status)
VALUES
  ("00","https://github.com/FhLk/SchoolSystem/blob/e3702ae0b5b72feaa14d35c34e88374d589f6105/Library-frontend/public/images/Mathematic.jpg",'คณิตศาสตร์', 'แบบฝึกหัดคณิตศาสตร์ ป.6', 'สถาบันส่งเสริมการสอนวิทยาศาสตร์และเทคโนโลยี', 'สสวท.', 2024, 0),
  ("01","https://github.com/FhLk/SchoolSystem/blob/e3702ae0b5b72feaa14d35c34e88374d589f6105/Library-frontend/public/images/Science.jpg",'วิทยาศาสตร์', 'วิทยาศาสตร์รอบตัว', 'กลุ่มนักเขียน', 'สำนักพิมพ์อมรินทร์', 2023, 0),
  ("02","https://github.com/FhLk/SchoolSystem/blob/e3702ae0b5b72feaa14d35c34e88374d589f6105/Library-frontend/public/images/Thai.jpg",'ภาษาไทย', 'ไวยากรณ์ภาษาไทย ม.3', 'อาจารย์สุภัทร สุวรรณ', 'พิงค์กีส์การศึกษา', 2022, 1);
  