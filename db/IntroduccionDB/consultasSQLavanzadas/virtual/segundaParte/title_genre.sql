SELECT m.title, g.name genre 
FROM movies m
INNER JOIN genres g
ON m.genre_id = g.id
