-- rambler up

ALTER TABLE `fortunes` RENAME TO `quotes`;

-- rambler down

ALTER TABLE `quotes` RENAME TO `fortunes`;

