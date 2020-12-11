--
-- Database: `go-crud`
--
CREATE DATABASE IF NOT EXISTS `go-crud` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `go-crud`;

-- --------------------------------------------------------

--
-- Table structure for table `mahasiswa`
--

CREATE TABLE `mahasiswa` (
  `nobp` varchar(100) NOT NULL PRIMARY KEY,
  `nama` varchar(100) NOT NULL,
  `alamat` varchar(100) NOT NULL,
  `umur` int(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
