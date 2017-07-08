set sql dialect 3;

CREATE TABLE r_selfconfig
(
    host_name       VARCHAR(128) NOT NULL,
    ipaddr          VARCHAR(24) NOT NULL,
    port            INT,	
    
    PRIMARY KEY (host_name)
);

CREATE TABLE r_adjnode
(
    nodeid           INT NOT NULL,
    nodetype         INT NOT NULL,
    ipaddr           VARCHAR(24) NOT NULL,
    port             INT,

    PRIMARY KEY (nodeid)
);

INSERT INTO r_selfconfig VALUES('ZTE-LBS', '10.43.31.148', 3001);

INSERT INTO r_adjnode VALUES(1, 0, '10.43.154.143', 4001);
INSERT INTO r_adjnode VALUES(2, 0, '10.43.154.143', 4002);


