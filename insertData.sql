CREATE  TABLE `library`.`books` (
  `Id` VARCHAR(255) NOT NULL,
  `Img` varchar(255),
  `Title` VARCHAR(45) NOT NULL,
  `Category` VARCHAR(45) NOT NULL,
  `Author` VARCHAR(45) NOT NULL,
  `Publisher` VARCHAR(45) NOT NULL,
  `Year` INT NOT NULL,
  `Status` INT NOT NULL,
  PRIMARY KEY (`Id`),
  UNIQUE INDEX `Id_UNIQUE` (`Id` ASC) VISIBLE);

drop table books;

INSERT INTO books (Id,Img, Category,Title, Author, Publisher, Year, Status)
VALUES
  ("00","https://github.com/FhLk/SchoolSystem/blob/e3702ae0b5b72feaa14d35c34e88374d589f6105/Library-frontend/public/images/Mathematic.jpg",'คณิตศาสตร์', 'แบบฝึกหัดคณิตศาสตร์ ป.6', 'สถาบันส่งเสริมการสอนวิทยาศาสตร์และเทคโนโลยี', 'สสวท.', 2024, 0),
  ("01","https://github.com/FhLk/SchoolSystem/blob/e3702ae0b5b72feaa14d35c34e88374d589f6105/Library-frontend/public/images/Science.jpg",'วิทยาศาสตร์', 'วิทยาศาสตร์รอบตัว', 'กลุ่มนักเขียน', 'สำนักพิมพ์อมรินทร์', 2023, 0),
  ("02","https://github.com/FhLk/SchoolSystem/blob/e3702ae0b5b72feaa14d35c34e88374d589f6105/Library-frontend/public/images/Thai.jpg",'ภาษาไทย', 'ไวยากรณ์ภาษาไทย ม.3', 'อาจารย์สุภัทร สุวรรณ', 'พิงค์กีส์การศึกษา', 2022, 1)
  ;
  