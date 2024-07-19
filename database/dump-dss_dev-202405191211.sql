--
-- PostgreSQL database dump
--

-- Dumped from database version 16.2
-- Dumped by pg_dump version 16.2

-- Started on 2024-05-19 12:11:38

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 5 (class 2615 OID 2200)
-- Name: dss_main; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA dss_main;


--
-- TOC entry 4830 (class 0 OID 0)
-- Dependencies: 5
-- Name: SCHEMA dss_main; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON SCHEMA dss_main IS 'standard public schema';


--
-- TOC entry 862 (class 1247 OID 16972)
-- Name: gender; Type: TYPE; Schema: dss_main; Owner: -
--

CREATE TYPE dss_main.gender AS ENUM (
    'L',
    'P'
);


--
-- TOC entry 850 (class 1247 OID 16411)
-- Name: yesno; Type: TYPE; Schema: dss_main; Owner: -
--

CREATE TYPE dss_main.yesno AS ENUM (
    'Y',
    'N'
);


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 219 (class 1259 OID 16488)
-- Name: tmp_mst_menu; Type: TABLE; Schema: dss_main; Owner: -
--

CREATE TABLE dss_main.tmp_mst_menu (
    menu_id integer NOT NULL,
    name character varying(55),
    class character varying(55),
    link character varying(255),
    level integer,
    parent integer,
    icon character varying(55),
    set_shortcut character varying,
    counter integer,
    description text,
    flag_menu character varying(55),
    is_active dss_main.yesno,
    created_at timestamp without time zone,
    created_by character varying(255),
    updated_at timestamp without time zone DEFAULT now(),
    updated_by character varying(255),
    display_on_tree dss_main.yesno DEFAULT 'Y'::dss_main.yesno,
    deleted_at timestamp without time zone
);


--
-- TOC entry 220 (class 1259 OID 16491)
-- Name: tmp_mst_menu_menu_id_seq; Type: SEQUENCE; Schema: dss_main; Owner: -
--

CREATE SEQUENCE dss_main.tmp_mst_menu_menu_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 4831 (class 0 OID 0)
-- Dependencies: 220
-- Name: tmp_mst_menu_menu_id_seq; Type: SEQUENCE OWNED BY; Schema: dss_main; Owner: -
--

ALTER SEQUENCE dss_main.tmp_mst_menu_menu_id_seq OWNED BY dss_main.tmp_mst_menu.menu_id;


--
-- TOC entry 217 (class 1259 OID 16427)
-- Name: tmp_mst_role; Type: TABLE; Schema: dss_main; Owner: -
--

CREATE TABLE dss_main.tmp_mst_role (
    role_id integer NOT NULL,
    level_id integer,
    name character varying(55),
    description text,
    is_active dss_main.yesno DEFAULT 'Y'::dss_main.yesno,
    created_at timestamp without time zone,
    created_by character varying(255),
    updated_at timestamp without time zone DEFAULT now(),
    updated_by character varying(255),
    deleted_at timestamp without time zone
);


--
-- TOC entry 218 (class 1259 OID 16430)
-- Name: tmp_mst_role_role_id_seq; Type: SEQUENCE; Schema: dss_main; Owner: -
--

CREATE SEQUENCE dss_main.tmp_mst_role_role_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 4832 (class 0 OID 0)
-- Dependencies: 218
-- Name: tmp_mst_role_role_id_seq; Type: SEQUENCE OWNED BY; Schema: dss_main; Owner: -
--

ALTER SEQUENCE dss_main.tmp_mst_role_role_id_seq OWNED BY dss_main.tmp_mst_role.role_id;


--
-- TOC entry 222 (class 1259 OID 16515)
-- Name: tmp_role_has_menu; Type: TABLE; Schema: dss_main; Owner: -
--

CREATE TABLE dss_main.tmp_role_has_menu (
    id integer NOT NULL,
    role_id integer,
    menu_id integer,
    action_code character varying(25)
);


--
-- TOC entry 221 (class 1259 OID 16514)
-- Name: tmp_role_has_menu_id_seq; Type: SEQUENCE; Schema: dss_main; Owner: -
--

CREATE SEQUENCE dss_main.tmp_role_has_menu_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 4833 (class 0 OID 0)
-- Dependencies: 221
-- Name: tmp_role_has_menu_id_seq; Type: SEQUENCE OWNED BY; Schema: dss_main; Owner: -
--

ALTER SEQUENCE dss_main.tmp_role_has_menu_id_seq OWNED BY dss_main.tmp_role_has_menu.id;


--
-- TOC entry 215 (class 1259 OID 16400)
-- Name: tmp_user; Type: TABLE; Schema: dss_main; Owner: -
--

CREATE TABLE dss_main.tmp_user (
    user_id integer NOT NULL,
    email character varying(255),
    fullname character varying(255),
    username character varying(255),
    password character varying(255),
    flag_user character varying(25),
    is_active dss_main.yesno,
    created_at timestamp without time zone,
    created_by character varying(255),
    updated_at timestamp without time zone DEFAULT now(),
    updated_by character varying(255),
    role_id integer,
    deleted_at timestamp without time zone
);


--
-- TOC entry 216 (class 1259 OID 16403)
-- Name: tmp_user_user_id_seq; Type: SEQUENCE; Schema: dss_main; Owner: -
--

CREATE SEQUENCE dss_main.tmp_user_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 4834 (class 0 OID 0)
-- Dependencies: 216
-- Name: tmp_user_user_id_seq; Type: SEQUENCE OWNED BY; Schema: dss_main; Owner: -
--

ALTER SEQUENCE dss_main.tmp_user_user_id_seq OWNED BY dss_main.tmp_user.user_id;


--
-- TOC entry 4660 (class 2604 OID 16492)
-- Name: tmp_mst_menu menu_id; Type: DEFAULT; Schema: dss_main; Owner: -
--

ALTER TABLE ONLY dss_main.tmp_mst_menu ALTER COLUMN menu_id SET DEFAULT nextval('dss_main.tmp_mst_menu_menu_id_seq'::regclass);


--
-- TOC entry 4657 (class 2604 OID 16431)
-- Name: tmp_mst_role role_id; Type: DEFAULT; Schema: dss_main; Owner: -
--

ALTER TABLE ONLY dss_main.tmp_mst_role ALTER COLUMN role_id SET DEFAULT nextval('dss_main.tmp_mst_role_role_id_seq'::regclass);


--
-- TOC entry 4663 (class 2604 OID 16518)
-- Name: tmp_role_has_menu id; Type: DEFAULT; Schema: dss_main; Owner: -
--

ALTER TABLE ONLY dss_main.tmp_role_has_menu ALTER COLUMN id SET DEFAULT nextval('dss_main.tmp_role_has_menu_id_seq'::regclass);


--
-- TOC entry 4655 (class 2604 OID 16404)
-- Name: tmp_user user_id; Type: DEFAULT; Schema: dss_main; Owner: -
--

ALTER TABLE ONLY dss_main.tmp_user ALTER COLUMN user_id SET DEFAULT nextval('dss_main.tmp_user_user_id_seq'::regclass);


--
-- TOC entry 4821 (class 0 OID 16488)
-- Dependencies: 219
-- Data for Name: tmp_mst_menu; Type: TABLE DATA; Schema: dss_main; Owner: -
--

INSERT INTO dss_main.tmp_mst_menu VALUES (4, 'Hak Akses Menu', NULL, '/settings/rolepreviledge', 0, 1, NULL, NULL, 3, NULL, NULL, 'Y', NULL, NULL, '2024-05-13 17:27:44.167291', NULL, 'Y', NULL);
INSERT INTO dss_main.tmp_mst_menu VALUES (3, 'Hak Akses', NULL, '/settings/roles', 0, 1, NULL, NULL, 2, NULL, NULL, 'Y', NULL, NULL, '2024-05-13 17:26:22.352297', NULL, 'Y', NULL);
INSERT INTO dss_main.tmp_mst_menu VALUES (5, 'Pengguna', NULL, '/settings/users', 0, 1, NULL, NULL, 4, NULL, NULL, 'Y', NULL, NULL, '2024-05-13 17:29:00.86352', NULL, 'Y', NULL);
INSERT INTO dss_main.tmp_mst_menu VALUES (1, 'Pengaturan', '#', '#', 0, 0, 'CogOutline', NULL, 0, NULL, NULL, 'Y', NULL, NULL, '2024-05-13 17:21:55.160378', NULL, 'Y', NULL);
INSERT INTO dss_main.tmp_mst_menu VALUES (2, 'Manajemen Menu', NULL, '/settings/menu', 0, 1, NULL, NULL, 1, NULL, NULL, 'Y', NULL, NULL, '2024-05-13 17:24:55.841488', NULL, 'Y', NULL);


--
-- TOC entry 4819 (class 0 OID 16427)
-- Dependencies: 217
-- Data for Name: tmp_mst_role; Type: TABLE DATA; Schema: dss_main; Owner: -
--

INSERT INTO dss_main.tmp_mst_role VALUES (1, 1, 'Admin Sistem', '', 'Y', '2017-12-18 21:43:15', '{"user_id":"1","fullname":"Amin Lubis"}', '2018-07-18 13:04:48', '{"user_id":"1","fullname":"Amin Lubis"}', NULL);


--
-- TOC entry 4824 (class 0 OID 16515)
-- Dependencies: 222
-- Data for Name: tmp_role_has_menu; Type: TABLE DATA; Schema: dss_main; Owner: -
--

INSERT INTO dss_main.tmp_role_has_menu VALUES (103, 1, 1, 'R');
INSERT INTO dss_main.tmp_role_has_menu VALUES (104, 1, 2, 'C');
INSERT INTO dss_main.tmp_role_has_menu VALUES (109, 1, 3, 'R');


--
-- TOC entry 4817 (class 0 OID 16400)
-- Dependencies: 215
-- Data for Name: tmp_user; Type: TABLE DATA; Schema: dss_main; Owner: -
--

INSERT INTO dss_main.tmp_user VALUES (1, 'deny.arwinto@gmail.com', 'Deny Arwinto', 'deny.arwinto', '$2a$08$5cP31y.AVVZEJbLNESekT.GReBSJdPB3jVBKp5bQVWGbk78.Nkhl6', NULL, NULL, NULL, NULL, '2024-05-10 10:25:33.22913', NULL, 1, NULL);


--
-- TOC entry 4835 (class 0 OID 0)
-- Dependencies: 220
-- Name: tmp_mst_menu_menu_id_seq; Type: SEQUENCE SET; Schema: dss_main; Owner: -
--

SELECT pg_catalog.setval('dss_main.tmp_mst_menu_menu_id_seq', 6, true);


--
-- TOC entry 4836 (class 0 OID 0)
-- Dependencies: 218
-- Name: tmp_mst_role_role_id_seq; Type: SEQUENCE SET; Schema: dss_main; Owner: -
--

SELECT pg_catalog.setval('dss_main.tmp_mst_role_role_id_seq', 1, false);


--
-- TOC entry 4837 (class 0 OID 0)
-- Dependencies: 221
-- Name: tmp_role_has_menu_id_seq; Type: SEQUENCE SET; Schema: dss_main; Owner: -
--

SELECT pg_catalog.setval('dss_main.tmp_role_has_menu_id_seq', 110, true);


--
-- TOC entry 4838 (class 0 OID 0)
-- Dependencies: 216
-- Name: tmp_user_user_id_seq; Type: SEQUENCE SET; Schema: dss_main; Owner: -
--

SELECT pg_catalog.setval('dss_main.tmp_user_user_id_seq', 1, true);


--
-- TOC entry 4669 (class 2606 OID 16553)
-- Name: tmp_mst_menu tmp_mst_menu_pk; Type: CONSTRAINT; Schema: dss_main; Owner: -
--

ALTER TABLE ONLY dss_main.tmp_mst_menu
    ADD CONSTRAINT tmp_mst_menu_pk PRIMARY KEY (menu_id);


--
-- TOC entry 4667 (class 2606 OID 16441)
-- Name: tmp_mst_role tmp_mst_role_pkey; Type: CONSTRAINT; Schema: dss_main; Owner: -
--

ALTER TABLE ONLY dss_main.tmp_mst_role
    ADD CONSTRAINT tmp_mst_role_pkey PRIMARY KEY (role_id);


--
-- TOC entry 4671 (class 2606 OID 16563)
-- Name: tmp_role_has_menu tmp_role_has_menu_pk; Type: CONSTRAINT; Schema: dss_main; Owner: -
--

ALTER TABLE ONLY dss_main.tmp_role_has_menu
    ADD CONSTRAINT tmp_role_has_menu_pk PRIMARY KEY (id);


--
-- TOC entry 4665 (class 2606 OID 16416)
-- Name: tmp_user tmp_user_pkey; Type: CONSTRAINT; Schema: dss_main; Owner: -
--

ALTER TABLE ONLY dss_main.tmp_user
    ADD CONSTRAINT tmp_user_pkey PRIMARY KEY (user_id);


--
-- TOC entry 4672 (class 2606 OID 16987)
-- Name: tmp_role_has_menu tmp_role_has_menu_fk1; Type: FK CONSTRAINT; Schema: dss_main; Owner: -
--

ALTER TABLE ONLY dss_main.tmp_role_has_menu
    ADD CONSTRAINT tmp_role_has_menu_fk1 FOREIGN KEY (role_id) REFERENCES dss_main.tmp_mst_role(role_id) ON DELETE CASCADE;


--
-- TOC entry 4673 (class 2606 OID 16992)
-- Name: tmp_role_has_menu tmp_role_has_menu_fk2; Type: FK CONSTRAINT; Schema: dss_main; Owner: -
--

ALTER TABLE ONLY dss_main.tmp_role_has_menu
    ADD CONSTRAINT tmp_role_has_menu_fk2 FOREIGN KEY (menu_id) REFERENCES dss_main.tmp_mst_menu(menu_id) ON DELETE CASCADE;


-- Completed on 2024-05-19 12:11:39

--
-- PostgreSQL database dump complete
--

