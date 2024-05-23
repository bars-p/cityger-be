INSERT INTO "role"(name, info) VALUES('admin', 'Power User');
INSERT INTO "role"(name, info) VALUES('user', 'Ordinary User');

INSERT INTO "client"(name, info) VALUES('Generic Client', 'Default client for general purpose');

INSERT INTO "user"(client_id, role_id, login, name)
    VALUES(1, 1, 'test', 'Generic test user');

INSERT INTO "project"(client_id, name, info, cell_area, zone_area, by_user_id)
    VALUES(1, 'Test project', 'Generic project for testing', 0.01, 1, 1);

INSERT INTO "scenario"(project_id, name, info, by_user_id)
    VALUES(1, 'Test scenario', 'Generic scenario for testing', 1);
