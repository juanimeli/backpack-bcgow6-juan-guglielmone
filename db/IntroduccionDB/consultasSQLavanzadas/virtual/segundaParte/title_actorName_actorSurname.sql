SELECT e.title episode_title, a.first_name, a.last_name
FROM episodes e
INNER JOIN actors a
INNER JOIN actor_episode ae
ON ae.actor_id = ae.episode_id