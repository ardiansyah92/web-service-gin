-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Mar 04, 2025 at 09:00 AM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.0.30

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_golang`
--

-- --------------------------------------------------------

--
-- Table structure for table `departements`
--

CREATE TABLE `departements` (
  `id` varchar(191) NOT NULL,
  `departement_name` longtext DEFAULT NULL,
  `location` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `files`
--

CREATE TABLE `files` (
  `id_file` bigint(20) NOT NULL,
  `filename` varchar(30) NOT NULL,
  `file_path` varchar(30) NOT NULL,
  `id_user` bigint(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `files`
--

INSERT INTO `files` (`id_file`, `filename`, `file_path`, `id_user`) VALUES
(1, '1741052848.png', './uploads/1741052848.png', 0),
(2, '1741053299.png', './uploads/1741053299.png', 0),
(3, '1741074815.png', './uploads/1741074815.png', 0),
(4, '1741074842.png', './uploads/1741074842.png', 0),
(5, '1741074953.png', './uploads/1741074953.png', 5);

-- --------------------------------------------------------

--
-- Table structure for table `loans`
--

CREATE TABLE `loans` (
  `id_loan` bigint(10) NOT NULL,
  `loan_application` int(11) NOT NULL,
  `interest_rate` varchar(10) NOT NULL,
  `month` varchar(20) NOT NULL,
  `user_loan` varchar(20) NOT NULL,
  `id_user` bigint(20) NOT NULL,
  `username` varchar(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `loans`
--

INSERT INTO `loans` (`id_loan`, `loan_application`, `interest_rate`, `month`, `user_loan`, `id_user`, `username`) VALUES
(1, 1000000, '1.0', '24', 'Rani', 5, 'rani');

-- --------------------------------------------------------

--
-- Stand-in structure for view `loan_view`
-- (See below for the actual view)
--
CREATE TABLE `loan_view` (
`pokok_pinjaman` double(17,0)
,`bunga_pertahun` double(17,0)
,`bunga_perbulan` double(17,0)
,`harus_dibayar` double
,`user` varchar(191)
);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id_user` bigint(20) NOT NULL,
  `username` varchar(191) NOT NULL,
  `password` longtext NOT NULL,
  `phone` varchar(191) NOT NULL,
  `email` varchar(191) NOT NULL,
  `address` varchar(191) NOT NULL,
  `user_loan` varchar(191) NOT NULL,
  `is_role` tinyint(1) DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id_user`, `username`, `password`, `phone`, `email`, `address`, `user_loan`, `is_role`) VALUES
(3, 'user', '$2a$10$9n5Rq.XNuQRzkFmrUaaF0.BlCAms.JTA33H74EBmswpRnlgefcmhu', '8912728812', 'user@yopmail.com', 'Jl.user SQL', 'user', 0),
(5, 'admin', '$2a$10$c8KXx/hnfiYYVKjmT5hfAuYeBGeDy9.wGv1L1szskATo7OgM8q0Pa', '089772881272', 'admin@yopmail.com', 'Jl.admin SQL', 'admin', 1),
(6, 'rani', '$2a$10$GCN8yk8xJhebWWblXRFrrui9BxE8YPGrStyuw3H6MdyICe6dQgmYa', '08977262712', 'rani@yopmail.com', 'Jl.Rani Runi', 'Rani', 0);

-- --------------------------------------------------------

--
-- Structure for view `loan_view`
--
DROP TABLE IF EXISTS `loan_view`;

CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `loan_view`  AS SELECT floor(`loans`.`loan_application` / `loans`.`month`) AS `pokok_pinjaman`, floor(`loans`.`loan_application` * `loans`.`interest_rate`) AS `bunga_pertahun`, floor(`loans`.`loan_application` * `loans`.`interest_rate` / `loans`.`month`) AS `bunga_perbulan`, floor(`loans`.`loan_application` / `loans`.`month`) + `loans`.`loan_application` * `loans`.`interest_rate` / `loans`.`month` AS `harus_dibayar`, `users`.`user_loan` AS `user` FROM (`loans` join `users` on(`loans`.`user_loan` = `users`.`user_loan`)) ;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `departements`
--
ALTER TABLE `departements`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `files`
--
ALTER TABLE `files`
  ADD PRIMARY KEY (`id_file`);

--
-- Indexes for table `loans`
--
ALTER TABLE `loans`
  ADD PRIMARY KEY (`id_loan`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id_user`),
  ADD UNIQUE KEY `uni_users_username` (`username`),
  ADD UNIQUE KEY `uni_users_phone` (`phone`),
  ADD UNIQUE KEY `uni_users_email` (`email`),
  ADD UNIQUE KEY `uni_users_address` (`address`),
  ADD UNIQUE KEY `uni_users_user_loan` (`user_loan`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `files`
--
ALTER TABLE `files`
  MODIFY `id_file` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `loans`
--
ALTER TABLE `loans`
  MODIFY `id_loan` bigint(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id_user` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
