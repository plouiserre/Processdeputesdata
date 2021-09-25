-- I use MySqlWorkBench which refuse to flush the table without where using a key column condition
DELETE FROM PROCESSDEPUTES.Mandate WHERE MandateId > 0;

DELETE FROM PROCESSDEPUTES.CongressMan WHERE CongressManId > 0;

DELETE FROM PROCESSDEPUTES.Deputy WHERE DeputyId > 0;