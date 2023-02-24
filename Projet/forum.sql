-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Hôte : 127.0.0.1:3306
-- Généré le : lun. 23 mai 2022 à 20:46
-- Version du serveur : 5.7.36
-- Version de PHP : 7.4.26

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de données : `forum`
--

-- --------------------------------------------------------

--
-- Structure de la table `messages`
--

DROP TABLE IF EXISTS `messages`;
CREATE TABLE IF NOT EXISTS `messages` (
  `id_message` int(10) NOT NULL AUTO_INCREMENT,
  `message` text NOT NULL,
  `date_message` varchar(254) NOT NULL,
  `id_forum` int(11) NOT NULL,
  `id_user` int(11) NOT NULL,
  PRIMARY KEY (`id_message`),
  KEY `id_forum` (`id_forum`)
) ENGINE=MyISAM AUTO_INCREMENT=20 DEFAULT CHARSET=latin1;

--
-- Déchargement des données de la table `messages`
--

INSERT INTO `messages` (`id_message`, `message`, `date_message`, `id_forum`, `id_user`) VALUES
(1, 'c\'est moi le plus bo', '2022-05-22', 21, 13),
(2, 'non c\'est faux', '2022-05-22', 21, 13),
(3, 'cc', '2022-05-22', 20, 13),
(4, 'ça doit être le poème de Maupassant, je l\'ai fait au collège ;)', '2022-05-23 22:05:19', 29, 17),
(5, 'Je te conseille le livre \"A la poursuite des sirènes\" de Galard, il résume bien les œuvres sur Ulysse', '2022-05-23 22:07:07', 26, 17),
(6, 'Généralement, c\'est marqué made in Japan derrière la peinture.', '2022-05-23 22:08:03', 30, 17),
(7, 'je me souvient plus mais la plupart ne meurent pas vraiment', '2022-05-23 22:11:08', 33, 17),
(8, 'Odin best forever', '2022-05-23 22:11:34', 35, 17),
(9, 'C\'est un rois à qui il lui arrive que des malheurs', '2022-05-23 22:12:15', 37, 17),
(10, 'C\'est un bonne question, up !', '2022-05-23 22:12:37', 39, 17),
(11, 'Normalement, il n\'en existe pas car c\'est le premier texte a introduire les origines de Zeus', '2022-05-23 22:20:05', 28, 16),
(12, 'On à longtemps pensés qu\'elle avait été fabriquée par des Étrusques, au Ve siècle av. J-C. Il apparaît qu\'elle date en fait du Moyen Âge^^', '2022-05-23 22:21:29', 27, 16),
(13, 'Beaucoup de personnes on fait des explications sur ce sujet mais les débats restent encore d\'actualités. Donc choisit ta voie jeune padawan.', '2022-05-23 22:23:02', 31, 16),
(14, 'Oui, Zeus à changé les souvenirs des habitants et à feinté leurs morts !', '2022-05-23 22:24:08', 33, 16),
(15, 'tu n\'as jamais vus Astérix xD ?', '2022-05-23 22:24:32', 34, 16),
(16, 'Tu bourrines pendant 3h, après ça passe mdr.', '2022-05-23 22:25:12', 36, 16),
(17, 'Je crois que c\'est dit dans un des textes d\'Hésiode', '2022-05-23 22:26:19', 39, 16),
(18, 'Oui effectivement, merci !', '2022-05-23 22:26:47', 29, 18),
(19, 'je l\'ai fait aussi !', '2022-05-23 22:27:40', 29, 16);

-- --------------------------------------------------------

--
-- Structure de la table `tabforum`
--

DROP TABLE IF EXISTS `tabforum`;
CREATE TABLE IF NOT EXISTS `tabforum` (
  `id_forum` int(10) NOT NULL AUTO_INCREMENT,
  `title_subject` varchar(254) NOT NULL,
  `categories` varchar(254) NOT NULL,
  `types` varchar(254) NOT NULL,
  `messages` text NOT NULL,
  `date` varchar(254) NOT NULL,
  `id_user` int(11) NOT NULL,
  PRIMARY KEY (`id_forum`)
) ENGINE=MyISAM AUTO_INCREMENT=42 DEFAULT CHARSET=latin1;

--
-- Déchargement des données de la table `tabforum`
--

INSERT INTO `tabforum` (`id_forum`, `title_subject`, `categories`, `types`, `messages`, `date`, `id_user`) VALUES
(29, 'Oedipe et l\'énigme du sphinx', 'Art', 'Egyptien', 'Je cherche un vieux texte lié à ce mythe mais je ne le retrouve plus, aidez moi svp !', '2022-05-23', 18),
(28, 'La Théogonie d\'Hésiode', 'Art', 'Grec', 'Connaissez vous des œuvres similaires ?', '2022-05-23', 18),
(26, 'Ulysse et les sirènes d\' Homère', 'Art', 'Grec', 'J\'adore l\'art et je voulais en apprendre plus ce cet œuvre', '2022-05-23', 18),
(27, 'La louve du Capitole', 'Art', 'Autre', 'D’où provient cet sculpture ?', '2022-05-23', 18),
(30, 'Reconnaitre les Yokais en peiture', 'Art', 'Japonais', 'Je possède un vase avec ce genre de motif, mais je ne m\'y connais pas trop, des suggestions ?', '2022-05-23', 18),
(31, 'Le serpent monde Jörmungand', 'Mythologie', 'Nordique', 'Quels sont vos théories sur la langue du serpent monde ?', '2022-05-23', 16),
(41, 'Mythe de Dédale', 'Mythologie', 'Grec', 'Vous le voyez plutôt comme un bon ou un mauvais personnage ?', '2022-05-23', 17),
(33, 'Le Ragnarôk', 'Mythologie', 'Nordique', 'Quels sont les personnages qui meurent pendant le Ragnarôk ?', '2022-05-23', 16),
(34, 'Les 12 Travaux d\'Ulysse', 'Mythologie', 'Grec', 'Pouvez vous répertorier les 12 travaux ?', '2022-05-23', 16),
(35, 'Les Dieux dans l\'Antiquité', 'Mythologie', 'Autre', 'Quel est votre tiers list des Dieux de l\'Antiquité ?', '2022-05-23', 16),
(36, 'Valkyries mythologie Nordique', 'Mythologie', 'Nordique', 'J\'arrive pas à les battres dans GOW4, des conseils svp ?', '2022-05-23', 17),
(37, 'Qui était le roi des héros Gilgamesh', 'Mythologie', 'Autre', 'Quel est l\'histoire de Gilgamesh ?', '2022-05-23', 17),
(38, 'Odin est-il mort ?', 'Mythologie', 'Nordique', 'On a pas vraiment la preuve qu\'il est mort, si ?', '2022-05-23', 17),
(39, 'Qui est la mère d\'Aphrodite', 'Mythologie', 'Grec', 'J\'attends vos théories sur le sujet:)', '2022-05-23', 17);

-- --------------------------------------------------------

--
-- Structure de la table `user`
--

DROP TABLE IF EXISTS `user`;
CREATE TABLE IF NOT EXISTS `user` (
  `id_user` int(10) NOT NULL AUTO_INCREMENT,
  `id_role` int(1) NOT NULL DEFAULT '0',
  `firstname` varchar(254) NOT NULL,
  `lastname` varchar(254) NOT NULL,
  `pseudo` varchar(254) NOT NULL,
  `born` varchar(254) NOT NULL,
  `email` varchar(254) NOT NULL,
  `password` varchar(254) NOT NULL,
  `lang` varchar(254) NOT NULL,
  `creation` varchar(254) NOT NULL,
  PRIMARY KEY (`id_user`),
  UNIQUE KEY `pseudo` (`pseudo`),
  UNIQUE KEY `email` (`email`)
) ENGINE=MyISAM AUTO_INCREMENT=21 DEFAULT CHARSET=latin1;

--
-- Déchargement des données de la table `user`
--

INSERT INTO `user` (`id_user`, `id_role`, `firstname`, `lastname`, `pseudo`, `born`, `email`, `password`, `lang`, `creation`) VALUES
(20, 1, 'Gabriel', 'Emmanuel', 'Gabriel_Emmanuel', '1995-01-01', 'gabriel@gmail.com', '63a9f0ea7bb98050796b649e85481845', 'Francais', '2022-05-23'),
(19, 1, 'Sébastien', 'Inzejob', 'Sébastien_Inzejob', '1992-01-01', 'sebastien@gmail.com', '63a9f0ea7bb98050796b649e85481845', 'Francais', '2022-05-23'),
(18, 0, 'Alexis', 'Giromagny', 'Rokuroo', '2000-03-26', 'alexis@gmail.com', '81dc9bdb52d04dc20036dbd8313ed055', 'Francais', '2022-05-23'),
(16, 0, 'gregory', 'balatre', 'Gregbal', '2003-07-30', 'gregory@gmail.com', '81dc9bdb52d04dc20036dbd8313ed055', 'Francais', '2022-05-23'),
(17, 0, 'Thomas', 'Quadro', 'Thetjqnhpp', '2003-09-11', 'thomas@gmail.com', '81dc9bdb52d04dc20036dbd8313ed055', 'Francais', '2022-05-23');
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
