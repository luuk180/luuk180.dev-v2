CREATE TABLE IF NOT EXISTS luuk180_dev.github_remotes (
    id text PRIMARY KEY NOT NULL ,
    name text,
    url text,
    homepageURL text,
    description text,
    diskusage int
)