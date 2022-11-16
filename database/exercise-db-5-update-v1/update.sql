-- TODO: answer here

-- WHERE ADDRESS NULL
UPDATE students SET address = 'Bandung'
WHERE address IS NULL AND status='active';
