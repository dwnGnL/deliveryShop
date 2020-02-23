-- phpMyAdmin SQL Dump
-- version 4.8.3
-- https://www.phpmyadmin.net/
--
-- Хост: 127.0.0.1:3306
-- Время создания: Фев 23 2020 г., 18:40
-- Версия сервера: 5.6.41
-- Версия PHP: 7.2.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- База данных: `deliveryfood`
--

-- --------------------------------------------------------

--
-- Структура таблицы `troles`
--

CREATE TABLE `troles` (
  `role_id` int(11) NOT NULL,
  `role_name` varchar(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Дамп данных таблицы `troles`
--

INSERT INTO `troles` (`role_id`, `role_name`) VALUES
(1, 'client'),
(2, 'admin');

-- --------------------------------------------------------

--
-- Структура таблицы `tusers`
--

CREATE TABLE `tusers` (
  `user_id` int(11) NOT NULL,
  `username` varchar(40) NOT NULL,
  `password` varchar(500) NOT NULL,
  `fullname` varchar(100) DEFAULT NULL,
  `phone` varchar(15) DEFAULT NULL,
  `salt` varchar(50) NOT NULL,
  `role` int(11) NOT NULL,
  `disabled` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Дамп данных таблицы `tusers`
--

INSERT INTO `tusers` (`user_id`, `username`, `password`, `fullname`, `phone`, `salt`, `role`, `disabled`) VALUES
(24, 'denzel', 'QYa73yClWR5eOh2Xu9TQePDRSeDNrQSmiOtEy2oBHcCLGgOOu4wL85m8M2mEBN8ns11mxkrJiG7/kxevSmqshA', 'denzel', '', 'RiHRr', 1, 0);

--
-- Индексы сохранённых таблиц
--

--
-- Индексы таблицы `troles`
--
ALTER TABLE `troles`
  ADD PRIMARY KEY (`role_id`);

--
-- Индексы таблицы `tusers`
--
ALTER TABLE `tusers`
  ADD PRIMARY KEY (`user_id`),
  ADD KEY `role` (`role`);

--
-- AUTO_INCREMENT для сохранённых таблиц
--

--
-- AUTO_INCREMENT для таблицы `troles`
--
ALTER TABLE `troles`
  MODIFY `role_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT для таблицы `tusers`
--
ALTER TABLE `tusers`
  MODIFY `user_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=25;

--
-- Ограничения внешнего ключа сохраненных таблиц
--

--
-- Ограничения внешнего ключа таблицы `tusers`
--
ALTER TABLE `tusers`
  ADD CONSTRAINT `tusers_ibfk_1` FOREIGN KEY (`role`) REFERENCES `troles` (`role_id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
