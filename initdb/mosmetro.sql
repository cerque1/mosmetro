--
-- PostgreSQL database dump
--

-- Dumped from database version 17.3
-- Dumped by pg_dump version 17.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: stations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.stations (
    id integer NOT NULL,
    name character varying NOT NULL,
    branch character varying NOT NULL
);


ALTER TABLE public.stations OWNER TO postgres;

--
-- Data for Name: stations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.stations (id, name, branch) FROM stdin;
1	Бульвар Дмитрия Донского	Серпуховско-Тимирязевская линия
2	Аннино	Серпуховско-Тимирязевская линия
3	Улица академика Янгеля	Серпуховско-Тимирязевская линия
4	Пражская	Серпуховско-Тимирязевская линия
5	Южная	Серпуховско-Тимирязевская линия
6	Чертановская	Серпуховско-Тимирязевская линия
7	Севастопольская	Серпуховско-Тимирязевская линия
8	Нахимовский Проспект	Серпуховско-Тимирязевская линия
9	Нагорная	Серпуховско-Тимирязевская линия
10	Нагатинская	Серпуховско-Тимирязевская линия
11	Тульская	Серпуховско-Тимирязевская линия
12	Серпуховская	Серпуховско-Тимирязевская линия
13	Полянка	Серпуховско-Тимирязевская линия
14	Боровицкая	Серпуховско-Тимирязевская линия
15	Чеховская	Серпуховско-Тимирязевская линия
16	Цветной бульвар	Серпуховско-Тимирязевская линия
17	Менделеевская	Серпуховско-Тимирязевская линия
18	Савёловская	Серпуховско-Тимирязевская линия
19	Дмитровская	Серпуховско-Тимирязевская линия
20	Тимирязевская	Серпуховско-Тимирязевская линия
21	Петровско-Разумовская	Серпуховско-Тимирязевская линия
22	Владыкино	Серпуховско-Тимирязевская линия
23	Отрадное	Серпуховско-Тимирязевская линия
24	Бибирево	Серпуховско-Тимирязевская линия
25	Алтуфьево	Серпуховско-Тимирязевская линия
\.


--
-- Name: stations stations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.stations
    ADD CONSTRAINT stations_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

