--
-- PostgreSQL database dump
--

-- Dumped from database version 16.2
-- Dumped by pg_dump version 16.2

-- Started on 2024-05-21 10:22:28

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
-- TOC entry 4832 (class 0 OID 0)
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
    parent integer,
    icon character varying(55),
    set_shortcut character varying,
    counter integer,
    description text,
    flag_menu character varying(55),
    is_active dss_main.yesno DEFAULT 'Y'::dss_main.yesno,
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
-- TOC entry 4833 (class 0 OID 0)
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
-- TOC entry 4834 (class 0 OID 0)
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
    action_code character varying(25) DEFAULT 'R'::character varying
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
-- TOC entry 4835 (class 0 OID 0)
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
-- TOC entry 4836 (class 0 OID 0)
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
-- TOC entry 4664 (class 2604 OID 16518)
-- Name: tmp_role_has_menu id; Type: DEFAULT; Schema: dss_main; Owner: -
--

ALTER TABLE ONLY dss_main.tmp_role_has_menu ALTER COLUMN id SET DEFAULT nextval('dss_main.tmp_role_has_menu_id_seq'::regclass);


--
-- TOC entry 4655 (class 2604 OID 16404)
-- Name: tmp_user user_id; Type: DEFAULT; Schema: dss_main; Owner: -
--

ALTER TABLE ONLY dss_main.tmp_user ALTER COLUMN user_id SET DEFAULT nextval('dss_main.tmp_user_user_id_seq'::regclass);


--
-- TOC entry 4823 (class 0 OID 16488)
-- Dependencies: 219
-- Data for Name: tmp_mst_menu; Type: TABLE DATA; Schema: dss_main; Owner: -
--

INSERT INTO dss_main.tmp_mst_menu VALUES (14, 'User List', NULL, '/master/userlist/', 13, NULL, NULL, 8, NULL, NULL, 'Y', NULL, NULL, '2024-05-21 10:18:51.596388', NULL, 'Y', NULL);
INSERT INTO dss_main.tmp_mst_menu VALUES (15, 'Daftarkan Fingerprint', NULL, '/master/fingerprint/', 13, NULL, NULL, 9, NULL, NULL, 'Y', NULL, NULL, '2024-05-21 10:19:31.875202', NULL, 'Y', NULL);
INSERT INTO dss_main.tmp_mst_menu VALUES (16, 'Setting Insya/Daily', NULL, '/master/inysadaily/', 6, NULL, NULL, 10, NULL, NULL, 'Y', NULL, NULL, '2024-05-21 10:20:47.124234', NULL, 'Y', NULL);
INSERT INTO dss_main.tmp_mst_menu VALUES (4, 'Hak Akses Menu', NULL, '/settings/rolepreviledge', 1, NULL, NULL, 3, NULL, NULL, 'Y', NULL, NULL, '2024-05-13 17:27:44.167291', NULL, 'Y', NULL);
INSERT INTO dss_main.tmp_mst_menu VALUES (3, 'Hak Akses', NULL, '/settings/roles', 1, NULL, NULL, 2, NULL, NULL, 'Y', NULL, NULL, '2024-05-13 17:26:22.352297', NULL, 'Y', NULL);
INSERT INTO dss_main.tmp_mst_menu VALUES (5, 'Pengguna', NULL, '/settings/user', 1, NULL, NULL, 4, NULL, NULL, 'Y', NULL, NULL, '2024-05-13 17:29:00.86352', NULL, 'Y', NULL);
INSERT INTO dss_main.tmp_mst_menu VALUES (2, 'Menu', NULL, '/settings/menu', 1, NULL, NULL, 1, NULL, NULL, 'Y', NULL, NULL, '2024-05-13 17:24:55.841488', NULL, 'Y', NULL);
INSERT INTO dss_main.tmp_mst_menu VALUES (7, 'Lembaga Pendidikan', NULL, '/master/cabang/', 6, NULL, NULL, 1, NULL, NULL, 'Y', NULL, NULL, '2024-05-21 09:01:12.543175', NULL, 'Y', NULL);
INSERT INTO dss_main.tmp_mst_menu VALUES (9, 'Asrama Rayon', NULL, '/master/asramarayon/', 8, NULL, NULL, 3, NULL, NULL, 'Y', NULL, NULL, '2024-05-21 09:07:25.47627', NULL, 'Y', NULL);
INSERT INTO dss_main.tmp_mst_menu VALUES (6, 'Master', '', '/master/', 0, 'BarsOutline', NULL, 0, NULL, NULL, 'Y', NULL, NULL, '2024-05-21 08:39:22.769437', NULL, 'Y', NULL);
INSERT INTO dss_main.tmp_mst_menu VALUES (1, 'Pengaturan', '#', '/settings/', 0, 'CogOutline', NULL, 100, NULL, NULL, 'Y', NULL, NULL, '2024-05-13 17:21:55.160378', NULL, 'Y', NULL);
INSERT INTO dss_main.tmp_mst_menu VALUES (10, 'List Kamar', NULL, '/master/listkamar/', 8, NULL, NULL, 4, NULL, NULL, 'Y', NULL, NULL, '2024-05-21 10:13:00.955761', NULL, 'Y', NULL);
INSERT INTO dss_main.tmp_mst_menu VALUES (11, 'Kamar Per Periode', NULL, '/master/kamarperiode/', 8, NULL, NULL, 5, NULL, NULL, 'Y', NULL, NULL, '2024-05-21 10:14:37.313675', NULL, 'Y', NULL);
INSERT INTO dss_main.tmp_mst_menu VALUES (12, 'Kamar Per Siswa', NULL, '/master/kamarsiswa/', 8, NULL, NULL, 6, NULL, NULL, 'Y', NULL, NULL, '2024-05-21 10:15:40.288328', NULL, 'Y', NULL);
INSERT INTO dss_main.tmp_mst_menu VALUES (13, 'User', NULL, '/master/user/', 6, NULL, NULL, 7, NULL, NULL, 'Y', NULL, NULL, '2024-05-21 10:17:49.471303', NULL, 'Y', NULL);
INSERT INTO dss_main.tmp_mst_menu VALUES (8, 'Rayon', NULL, '/master/rayon/', 6, NULL, NULL, 2, NULL, NULL, 'Y', NULL, NULL, '2024-05-21 09:06:35.530504', NULL, 'Y', NULL);


--
-- TOC entry 4821 (class 0 OID 16427)
-- Dependencies: 217
-- Data for Name: tmp_mst_role; Type: TABLE DATA; Schema: dss_main; Owner: -
--

INSERT INTO dss_main.tmp_mst_role VALUES (1, 1, 'Super Admin', '', 'Y', '2017-12-18 21:43:15', '{"user_id":"1","fullname":"Amin Lubis"}', '2018-07-18 13:04:48', '{"user_id":"1","fullname":"Amin Lubis"}', NULL);


--
-- TOC entry 4826 (class 0 OID 16515)
-- Dependencies: 222
-- Data for Name: tmp_role_has_menu; Type: TABLE DATA; Schema: dss_main; Owner: -
--

INSERT INTO dss_main.tmp_role_has_menu VALUES (103, 1, 1, 'R');
INSERT INTO dss_main.tmp_role_has_menu VALUES (109, 1, 3, 'R');
INSERT INTO dss_main.tmp_role_has_menu VALUES (111, 1, 5, 'R');
INSERT INTO dss_main.tmp_role_has_menu VALUES (104, 1, 2, 'R');
INSERT INTO dss_main.tmp_role_has_menu VALUES (112, 1, 6, 'R');
INSERT INTO dss_main.tmp_role_has_menu VALUES (113, 1, 7, 'R');
INSERT INTO dss_main.tmp_role_has_menu VALUES (114, 1, 8, 'R');
INSERT INTO dss_main.tmp_role_has_menu VALUES (115, 1, 9, 'R');
INSERT INTO dss_main.tmp_role_has_menu VALUES (116, 1, 10, 'R');
INSERT INTO dss_main.tmp_role_has_menu VALUES (117, 1, 11, 'R');
INSERT INTO dss_main.tmp_role_has_menu VALUES (118, 1, 12, 'R');
INSERT INTO dss_main.tmp_role_has_menu VALUES (119, 1, 13, 'R');
INSERT INTO dss_main.tmp_role_has_menu VALUES (120, 1, 14, 'R');
INSERT INTO dss_main.tmp_role_has_menu VALUES (121, 1, 15, 'R');
INSERT INTO dss_main.tmp_role_has_menu VALUES (122, 1, 16, 'R');


--
-- TOC entry 4819 (class 0 OID 16400)
-- Dependencies: 215
-- Data for Name: tmp_user; Type: TABLE DATA; Schema: dss_main; Owner: -
--

INSERT INTO dss_main.tmp_user VALUES (1, 'deny.arwinto@gmail.com', 'Deny Arwinto', 'deny.arwinto', '$2a$08$5cP31y.AVVZEJbLNESekT.GReBSJdPB3jVBKp5bQVWGbk78.Nkhl6', NULL, 'Y', NULL, NULL, '2024-05-10 10:25:33.22913', NULL, 1, NULL);
INSERT INTO dss_main.tmp_user VALUES (2, 'deny.arwinto@gmail.com', 'Deddy Kristanto', NULL, NULL, NULL, 'Y', NULL, NULL, '2024-05-21 08:26:43.822499', NULL, 1, NULL);
INSERT INTO dss_main.tmp_user VALUES (3, 'deny.arwinto@gmail.com', 'Admin Wisata Alam Endah', NULL, NULL, NULL, 'Y', NULL, NULL, '2024-05-21 08:29:57.300531', NULL, 1, NULL);


--
-- TOC entry 4837 (class 0 OID 0)
-- Dependencies: 220
-- Name: tmp_mst_menu_menu_id_seq; Type: SEQUENCE SET; Schema: dss_main; Owner: -
--

SELECT pg_catalog.setval('dss_main.tmp_mst_menu_menu_id_seq', 16, true);


--
-- TOC entry 4838 (class 0 OID 0)
-- Dependencies: 218
-- Name: tmp_mst_role_role_id_seq; Type: SEQUENCE SET; Schema: dss_main; Owner: -
--

SELECT pg_catalog.setval('dss_main.tmp_mst_role_role_id_seq', 1, false);


--
-- TOC entry 4839 (class 0 OID 0)
-- Dependencies: 221
-- Name: tmp_role_has_menu_id_seq; Type: SEQUENCE SET; Schema: dss_main; Owner: -
--

SELECT pg_catalog.setval('dss_main.tmp_role_has_menu_id_seq', 122, true);


--
-- TOC entry 4840 (class 0 OID 0)
-- Dependencies: 216
-- Name: tmp_user_user_id_seq; Type: SEQUENCE SET; Schema: dss_main; Owner: -
--

SELECT pg_catalog.setval('dss_main.tmp_user_user_id_seq', 3, true);


--
-- TOC entry 4671 (class 2606 OID 16553)
-- Name: tmp_mst_menu tmp_mst_menu_pk; Type: CONSTRAINT; Schema: dss_main; Owner: -
--

ALTER TABLE ONLY dss_main.tmp_mst_menu
    ADD CONSTRAINT tmp_mst_menu_pk PRIMARY KEY (menu_id);


--
-- TOC entry 4669 (class 2606 OID 16441)
-- Name: tmp_mst_role tmp_mst_role_pkey; Type: CONSTRAINT; Schema: dss_main; Owner: -
--

ALTER TABLE ONLY dss_main.tmp_mst_role
    ADD CONSTRAINT tmp_mst_role_pkey PRIMARY KEY (role_id);


--
-- TOC entry 4673 (class 2606 OID 16563)
-- Name: tmp_role_has_menu tmp_role_has_menu_pk; Type: CONSTRAINT; Schema: dss_main; Owner: -
--

ALTER TABLE ONLY dss_main.tmp_role_has_menu
    ADD CONSTRAINT tmp_role_has_menu_pk PRIMARY KEY (id);


--
-- TOC entry 4667 (class 2606 OID 16416)
-- Name: tmp_user tmp_user_pkey; Type: CONSTRAINT; Schema: dss_main; Owner: -
--

ALTER TABLE ONLY dss_main.tmp_user
    ADD CONSTRAINT tmp_user_pkey PRIMARY KEY (user_id);


--
-- TOC entry 4674 (class 2606 OID 16987)
-- Name: tmp_role_has_menu tmp_role_has_menu_fk1; Type: FK CONSTRAINT; Schema: dss_main; Owner: -
--

ALTER TABLE ONLY dss_main.tmp_role_has_menu
    ADD CONSTRAINT tmp_role_has_menu_fk1 FOREIGN KEY (role_id) REFERENCES dss_main.tmp_mst_role(role_id) ON DELETE CASCADE;


--
-- TOC entry 4675 (class 2606 OID 16992)
-- Name: tmp_role_has_menu tmp_role_has_menu_fk2; Type: FK CONSTRAINT; Schema: dss_main; Owner: -
--

ALTER TABLE ONLY dss_main.tmp_role_has_menu
    ADD CONSTRAINT tmp_role_has_menu_fk2 FOREIGN KEY (menu_id) REFERENCES dss_main.tmp_mst_menu(menu_id) ON DELETE CASCADE;


-- Completed on 2024-05-21 10:22:29

--
-- PostgreSQL database dump complete
--

