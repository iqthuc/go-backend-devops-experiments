--
-- PostgreSQL database dump
--

-- Dumped from database version 17.4 (Debian 17.4-1.pgdg120+2)
-- Dumped by pg_dump version 17.2

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
-- Name: brands; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.brands (
    id integer NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE public.brands OWNER TO root;

--
-- Name: brands_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.brands_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.brands_id_seq OWNER TO root;

--
-- Name: brands_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.brands_id_seq OWNED BY public.brands.id;


--
-- Name: categories; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.categories (
    id integer NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE public.categories OWNER TO root;

--
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.categories_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.categories_id_seq OWNER TO root;

--
-- Name: categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;


--
-- Name: colors; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.colors (
    id integer NOT NULL,
    name character varying(32) NOT NULL
);


ALTER TABLE public.colors OWNER TO root;

--
-- Name: colors_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.colors_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.colors_id_seq OWNER TO root;

--
-- Name: colors_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.colors_id_seq OWNED BY public.colors.id;


--
-- Name: product_details; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.product_details (
    id integer NOT NULL,
    product_id integer,
    description text,
    technical_info jsonb,
    images jsonb
);


ALTER TABLE public.product_details OWNER TO root;

--
-- Name: product_details_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.product_details_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.product_details_id_seq OWNER TO root;

--
-- Name: product_details_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.product_details_id_seq OWNED BY public.product_details.id;


--
-- Name: product_discounts; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.product_discounts (
    id integer NOT NULL,
    product_variant_id integer,
    discount_price numeric(10,2) NOT NULL,
    start_date date NOT NULL,
    end_date date NOT NULL
);


ALTER TABLE public.product_discounts OWNER TO root;

--
-- Name: product_discounts_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.product_discounts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.product_discounts_id_seq OWNER TO root;

--
-- Name: product_discounts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.product_discounts_id_seq OWNED BY public.product_discounts.id;


--
-- Name: product_images; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.product_images (
    id integer NOT NULL,
    product_id integer,
    images_url character varying NOT NULL,
    is_primary boolean DEFAULT false
);


ALTER TABLE public.product_images OWNER TO root;

--
-- Name: product_images_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.product_images_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.product_images_id_seq OWNER TO root;

--
-- Name: product_images_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.product_images_id_seq OWNED BY public.product_images.id;


--
-- Name: product_prices; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.product_prices (
    id integer NOT NULL,
    product__variant_id integer,
    price numeric(10,2) NOT NULL
);


ALTER TABLE public.product_prices OWNER TO root;

--
-- Name: product_prices_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.product_prices_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.product_prices_id_seq OWNER TO root;

--
-- Name: product_prices_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.product_prices_id_seq OWNED BY public.product_prices.id;


--
-- Name: product_ratings; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.product_ratings (
    id integer NOT NULL,
    product_id integer,
    rating integer NOT NULL,
    comment text,
    created_at timestamp with time zone DEFAULT now()
);


ALTER TABLE public.product_ratings OWNER TO root;

--
-- Name: product_ratings_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.product_ratings_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.product_ratings_id_seq OWNER TO root;

--
-- Name: product_ratings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.product_ratings_id_seq OWNED BY public.product_ratings.id;


--
-- Name: product_variant_colors; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.product_variant_colors (
    id integer NOT NULL,
    product_variant_id integer,
    color_id integer
);


ALTER TABLE public.product_variant_colors OWNER TO root;

--
-- Name: product_variant_colors_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.product_variant_colors_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.product_variant_colors_id_seq OWNER TO root;

--
-- Name: product_variant_colors_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.product_variant_colors_id_seq OWNED BY public.product_variant_colors.id;


--
-- Name: product_variant_sizes; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.product_variant_sizes (
    id integer NOT NULL,
    product_variant_id integer,
    size_id integer
);


ALTER TABLE public.product_variant_sizes OWNER TO root;

--
-- Name: product_variant_sizes_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.product_variant_sizes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.product_variant_sizes_id_seq OWNER TO root;

--
-- Name: product_variant_sizes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.product_variant_sizes_id_seq OWNED BY public.product_variant_sizes.id;


--
-- Name: product_variants; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.product_variants (
    id integer NOT NULL,
    product_id integer,
    stock_quantity integer DEFAULT 0
);


ALTER TABLE public.product_variants OWNER TO root;

--
-- Name: product_variants_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.product_variants_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.product_variants_id_seq OWNER TO root;

--
-- Name: product_variants_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.product_variants_id_seq OWNED BY public.product_variants.id;


--
-- Name: products; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.products (
    id integer NOT NULL,
    name character varying NOT NULL,
    category_id integer,
    brand_id integer
);


ALTER TABLE public.products OWNER TO root;

--
-- Name: products_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.products_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.products_id_seq OWNER TO root;

--
-- Name: products_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.products_id_seq OWNED BY public.products.id;


--
-- Name: sizes; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.sizes (
    id integer NOT NULL,
    name character varying(16) NOT NULL
);


ALTER TABLE public.sizes OWNER TO root;

--
-- Name: sizes_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.sizes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sizes_id_seq OWNER TO root;

--
-- Name: sizes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.sizes_id_seq OWNED BY public.sizes.id;


--
-- Name: brands id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.brands ALTER COLUMN id SET DEFAULT nextval('public.brands_id_seq'::regclass);


--
-- Name: categories id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.categories ALTER COLUMN id SET DEFAULT nextval('public.categories_id_seq'::regclass);


--
-- Name: colors id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.colors ALTER COLUMN id SET DEFAULT nextval('public.colors_id_seq'::regclass);


--
-- Name: product_details id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.product_details ALTER COLUMN id SET DEFAULT nextval('public.product_details_id_seq'::regclass);


--
-- Name: product_discounts id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.product_discounts ALTER COLUMN id SET DEFAULT nextval('public.product_discounts_id_seq'::regclass);


--
-- Name: product_images id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.product_images ALTER COLUMN id SET DEFAULT nextval('public.product_images_id_seq'::regclass);


--
-- Name: product_prices id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.product_prices ALTER COLUMN id SET DEFAULT nextval('public.product_prices_id_seq'::regclass);


--
-- Name: product_ratings id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.product_ratings ALTER COLUMN id SET DEFAULT nextval('public.product_ratings_id_seq'::regclass);


--
-- Name: product_variant_colors id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.product_variant_colors ALTER COLUMN id SET DEFAULT nextval('public.product_variant_colors_id_seq'::regclass);


--
-- Name: product_variant_sizes id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.product_variant_sizes ALTER COLUMN id SET DEFAULT nextval('public.product_variant_sizes_id_seq'::regclass);


--
-- Name: product_variants id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.product_variants ALTER COLUMN id SET DEFAULT nextval('public.product_variants_id_seq'::regclass);


--
-- Name: products id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.products ALTER COLUMN id SET DEFAULT nextval('public.products_id_seq'::regclass);


--
-- Name: sizes id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.sizes ALTER COLUMN id SET DEFAULT nextval('public.sizes_id_seq'::regclass);


--
-- Data for Name: brands; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.brands (id, name) FROM stdin;
1	Yonex
2	Victor
3	Lining
4	Mizuno
5	Apacs
6	Proace
\.


--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.categories (id, name) FROM stdin;
1	Vợt cầu lông
2	Áo cầu lông
3	Quần cầu lông
4	Giày cầu lông
5	Phụ kiện
\.


--
-- Data for Name: colors; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.colors (id, name) FROM stdin;
1	Đen
2	Trắng
3	Đỏ
4	Xanh dương
5	Vàng
6	Xanh lá cây
7	Cam
8	Tím
9	Hồng
10	Xám
11	Nâu
12	Xanh ngọc
\.


--
-- Data for Name: product_details; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.product_details (id, product_id, description, technical_info, images) FROM stdin;
\.


--
-- Data for Name: product_discounts; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.product_discounts (id, product_variant_id, discount_price, start_date, end_date) FROM stdin;
\.


--
-- Data for Name: product_images; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.product_images (id, product_id, images_url, is_primary) FROM stdin;
1	1	images/vot_cau_long/yonex_astrox_100zz_den_1.jpg	t
2	1	images/vot_cau_long/yonex_astrox_100zz_do_1.jpg	f
3	1	images/vot_cau_long/yonex_astrox_100zz_xanhduong_1.jpg	f
4	1	images/vot_cau_long/yonex_astrox_100zz_vang_1.jpg	f
5	1	images/vot_cau_long/yonex_astrox_100zz_xanhla_1.jpg	f
6	11	images/ao_cau_long/yonex_10452ex_den_1.jpg	t
7	11	images/ao_cau_long/yonex_10452ex_trang_1.jpg	f
8	11	images/ao_cau_long/yonex_10452ex_do_1.jpg	f
9	11	images/ao_cau_long/yonex_10452ex_xanhduong_1.jpg	f
10	11	images/ao_cau_long/yonex_10452ex_vang_1.jpg	f
11	21	images/quan_cau_long/yonex_15098ex_den_1.jpg	t
12	21	images/quan_cau_long/yonex_15098ex_trang_1.jpg	f
13	21	images/quan_cau_long/yonex_15098ex_xam_1.jpg	f
14	21	images/quan_cau_long/yonex_15098ex_xanhduong_1.jpg	f
15	21	images/quan_cau_long/yonex_15098ex_nau_1.jpg	f
16	31	images/giay_cau_long/yonex_aerus_z2_trang_1.jpg	t
17	31	images/giay_cau_long/yonex_aerus_z2_den_1.jpg	f
18	31	images/giay_cau_long/yonex_aerus_z2_xanhduong_1.jpg	f
19	31	images/giay_cau_long/yonex_aerus_z2_do_1.jpg	f
20	31	images/giay_cau_long/yonex_aerus_z2_xam_1.jpg	f
21	41	images/phu_kien_cau_long/yonex_ac102ex_den_1.jpg	t
22	41	images/phu_kien_cau_long/yonex_ac102ex_trang_1.jpg	f
23	41	images/phu_kien_cau_long/yonex_ac102ex_vang_1.jpg	f
24	41	images/phu_kien_cau_long/yonex_ac102ex_xanhduong_1.jpg	f
25	41	images/phu_kien_cau_long/yonex_ac102ex_do_1.jpg	f
\.


--
-- Data for Name: product_prices; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.product_prices (id, product__variant_id, price) FROM stdin;
1	1	3500000.00
2	2	3600000.00
3	3	3700000.00
4	4	3800000.00
5	5	3900000.00
6	6	3400000.00
7	7	3300000.00
8	8	3200000.00
9	9	3100000.00
10	10	3000000.00
11	51	500000.00
12	52	550000.00
13	53	600000.00
14	54	650000.00
15	55	700000.00
16	56	450000.00
17	57	400000.00
18	58	350000.00
19	59	300000.00
20	60	250000.00
21	66	400000.00
22	67	450000.00
23	68	500000.00
24	69	550000.00
25	70	600000.00
26	71	350000.00
27	72	300000.00
28	73	250000.00
29	74	200000.00
30	75	150000.00
31	81	2500000.00
32	82	2600000.00
33	83	2700000.00
34	84	2800000.00
35	85	2900000.00
36	86	2400000.00
37	87	2300000.00
38	88	2200000.00
39	89	2100000.00
40	90	2000000.00
41	96	50000.00
42	97	60000.00
43	98	70000.00
44	99	80000.00
45	100	90000.00
46	101	40000.00
47	102	30000.00
48	103	20000.00
49	104	10000.00
50	105	5000.00
\.


--
-- Data for Name: product_ratings; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.product_ratings (id, product_id, rating, comment, created_at) FROM stdin;
\.


--
-- Data for Name: product_variant_colors; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.product_variant_colors (id, product_variant_id, color_id) FROM stdin;
85	96	1
86	97	2
87	98	5
88	99	4
89	100	3
\.


--
-- Data for Name: product_variant_sizes; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.product_variant_sizes (id, product_variant_id, size_id) FROM stdin;
\.


--
-- Data for Name: product_variants; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.product_variants (id, product_id, stock_quantity) FROM stdin;
1	1	10
2	1	15
3	1	20
4	1	12
5	1	18
6	1	25
7	1	8
8	1	30
9	1	14
10	1	22
11	2	11
12	2	16
13	2	21
14	2	13
15	2	19
16	2	26
17	2	9
18	2	31
19	2	15
20	2	23
21	3	12
22	3	17
23	3	22
24	3	14
25	3	20
26	3	27
27	3	10
28	3	32
29	3	16
30	3	24
31	4	13
32	4	18
33	4	23
34	4	15
35	4	21
36	4	28
37	4	11
38	4	33
39	4	17
40	4	25
41	5	14
42	5	19
43	5	24
44	5	16
45	5	22
46	5	29
47	5	12
48	5	34
49	5	18
50	5	26
51	6	15
52	6	20
53	6	12
54	7	18
55	7	25
56	7	8
57	7	30
58	8	14
59	8	22
60	8	11
61	9	16
62	9	21
63	9	13
64	9	19
65	10	26
66	10	9
67	10	31
68	10	15
69	10	23
70	11	100
71	11	50
72	11	200
73	12	150
74	12	100
75	13	120
76	13	80
77	13	180
78	14	90
79	14	110
80	16	140
81	16	60
82	17	170
83	17	30
84	17	190
85	19	200
86	19	10
87	19	220
88	20	210
89	20	5
90	41	100
91	41	50
92	41	200
93	42	150
94	42	100
95	43	120
96	43	80
97	43	180
98	44	90
99	44	110
100	45	130
101	45	70
102	45	160
103	46	140
104	46	60
105	47	170
106	47	30
107	47	190
108	48	180
109	48	20
110	49	200
111	49	10
112	49	220
113	50	210
114	50	5
\.


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.products (id, name, category_id, brand_id) FROM stdin;
1	Yonex Astrox 100 ZZ	1	1
2	Victor Thruster F Claw LTD	1	2
3	Lining Aeronaut 9000C	1	3
4	Mizuno Fortius 30 Power	1	4
5	Apacs Virtuoso Performance	1	5
6	Proace Stroke 318 III	1	6
7	Yonex Arcsaber 11 Pro	1	1
8	Victor Brave Sword 12	1	2
9	Lining Turbo Charging 75	1	3
10	Mizuno Caliber VS Tour	1	4
11	Yonex 10452EX	2	1
12	Victor T-2000	2	2
13	Lining AAYQ157-1	2	3
14	Mizuno 72MA9001	2	4
15	Apacs AP-TT	2	5
16	Proace PR-AT	2	6
17	Yonex 10452EX - Xanh dương	2	1
18	Victor T-2000 - Đỏ	2	2
19	Lining AAYQ157-1 - Trắng	2	3
20	Mizuno 72MA9001 - Đen	2	4
21	Yonex 15098EX	3	1
22	Victor R-0900	3	2
23	Lining AKSP117-1	3	3
24	Mizuno 72MB9001	3	4
25	Apacs AP-PT	3	5
26	Proace PR-PT	3	6
27	Yonex 15098EX - Đen	3	1
28	Victor R-0900 - Xanh dương	3	2
29	Lining AKSP117-1 - Trắng	3	3
30	Mizuno 72MB9001 - Xám	3	4
31	Yonex Aerus Z2	4	1
32	Victor A970ACE	4	2
33	Lining AYAR011-1	4	3
34	Mizuno Wave Fang Pro	4	4
35	Apacs Lethal Light	4	5
36	Proace Power Ace	4	6
37	Yonex Aerus Z2 - Trắng	4	1
38	Victor A970ACE - Đen	4	2
39	Lining AYAR011-1 - Xanh dương	4	3
40	Mizuno Wave Fang Pro - Đỏ	4	4
41	Yonex AC102EX Grip	5	1
42	Victor GR262 Grip	5	2
43	Lining GP1000 Grip	5	3
44	Mizuno 63JYA801 Grip	5	4
45	Apacs Grip PU-12	5	5
46	Proace Grip PR-G	5	6
47	Yonex AS-50 Shuttlecock	5	1
48	Victor Master No.1 Shuttlecock	5	2
49	Lining A+90 Shuttlecock	5	3
50	Mizuno SKYCROSS R-1 Shuttlecock	5	4
\.


--
-- Data for Name: sizes; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.sizes (id, name) FROM stdin;
1	S
2	M
3	L
4	XL
5	40
6	41
7	42
8	43
9	44
10	45
11	46
12	47
13	48
14	49
15	50
\.


--
-- Name: brands_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.brands_id_seq', 6, true);


--
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.categories_id_seq', 5, true);


--
-- Name: colors_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.colors_id_seq', 12, true);


--
-- Name: product_details_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.product_details_id_seq', 1, false);


--
-- Name: product_discounts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.product_discounts_id_seq', 1, false);


--
-- Name: product_images_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.product_images_id_seq', 25, true);


--
-- Name: product_prices_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.product_prices_id_seq', 50, true);


--
-- Name: product_ratings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.product_ratings_id_seq', 1, false);


--
-- Name: product_variant_colors_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.product_variant_colors_id_seq', 89, true);


--
-- Name: product_variant_sizes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.product_variant_sizes_id_seq', 1, false);


--
-- Name: product_variants_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.product_variants_id_seq', 114, true);


--
-- Name: products_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.products_id_seq', 50, true);


--
-- Name: sizes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.sizes_id_seq', 15, true);


--
-- Name: brands brands_name_key; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.brands
    ADD CONSTRAINT brands_name_key UNIQUE (name);


--
-- Name: brands brands_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.brands
    ADD CONSTRAINT brands_pkey PRIMARY KEY (id);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- Name: colors colors_name_key; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.colors
    ADD CONSTRAINT colors_name_key UNIQUE (name);


--
-- Name: colors colors_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.colors
    ADD CONSTRAINT colors_pkey PRIMARY KEY (id);


--
-- Name: product_details product_details_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.product_details
    ADD CONSTRAINT product_details_pkey PRIMARY KEY (id);


--
-- Name: product_details product_details_product_id_key; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.product_details
    ADD CONSTRAINT product_details_product_id_key UNIQUE (product_id);


--
-- Name: product_discounts product_discounts_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.product_discounts
    ADD CONSTRAINT product_discounts_pkey PRIMARY KEY (id);


--
-- Name: product_images product_images_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.product_images
    ADD CONSTRAINT product_images_pkey PRIMARY KEY (id);


--
-- Name: product_prices product_prices_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.product_prices
    ADD CONSTRAINT product_prices_pkey PRIMARY KEY (id);


--
-- Name: product_ratings product_ratings_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.product_ratings
    ADD CONSTRAINT product_ratings_pkey PRIMARY KEY (id);


--
-- Name: product_variant_colors product_variant_colors_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.product_variant_colors
    ADD CONSTRAINT product_variant_colors_pkey PRIMARY KEY (id);


--
-- Name: product_variant_sizes product_variant_sizes_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.product_variant_sizes
    ADD CONSTRAINT product_variant_sizes_pkey PRIMARY KEY (id);


--
-- Name: product_variants product_variants_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.product_variants
    ADD CONSTRAINT product_variants_pkey PRIMARY KEY (id);


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- Name: sizes sizes_name_key; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.sizes
    ADD CONSTRAINT sizes_name_key UNIQUE (name);


--
-- Name: sizes sizes_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.sizes
    ADD CONSTRAINT sizes_pkey PRIMARY KEY (id);


--
-- Name: product_discounts_discount_price_idx; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX product_discounts_discount_price_idx ON public.product_discounts USING btree (discount_price);


--
-- Name: product_discounts_end_date_idx; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX product_discounts_end_date_idx ON public.product_discounts USING btree (end_date);


--
-- Name: product_discounts_product_variant_id_idx; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX product_discounts_product_variant_id_idx ON public.product_discounts USING btree (product_variant_id);


--
-- Name: product_discounts_start_date_idx; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX product_discounts_start_date_idx ON public.product_discounts USING btree (start_date);


--
-- Name: product_images_product_id_idx; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX product_images_product_id_idx ON public.product_images USING btree (product_id);


--
-- Name: product_prices_product__variant_id_idx; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX product_prices_product__variant_id_idx ON public.product_prices USING btree (product__variant_id);


--
-- Name: product_ratings_product_id_idx; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX product_ratings_product_id_idx ON public.product_ratings USING btree (product_id);


--
-- Name: product_variant_colors_color_id_idx; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX product_variant_colors_color_id_idx ON public.product_variant_colors USING btree (color_id);


--
-- Name: product_variant_colors_product_variant_id_idx; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX product_variant_colors_product_variant_id_idx ON public.product_variant_colors USING btree (product_variant_id);


--
-- Name: product_variant_sizes_product_variant_id_idx; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX product_variant_sizes_product_variant_id_idx ON public.product_variant_sizes USING btree (product_variant_id);


--
-- Name: product_variant_sizes_size_id_idx; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX product_variant_sizes_size_id_idx ON public.product_variant_sizes USING btree (size_id);


--
-- Name: product_variants_product_id_idx; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX product_variants_product_id_idx ON public.product_variants USING btree (product_id);


--
-- Name: products_category_id_idx; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX products_category_id_idx ON public.products USING btree (category_id);


--
-- Name: product_details product_details_product_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.product_details
    ADD CONSTRAINT product_details_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id);


--
-- Name: product_discounts product_discounts_product_variant_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.product_discounts
    ADD CONSTRAINT product_discounts_product_variant_id_fkey FOREIGN KEY (product_variant_id) REFERENCES public.product_variants(id);


--
-- Name: product_images product_images_product_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.product_images
    ADD CONSTRAINT product_images_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.product_variants(id);


--
-- Name: product_prices product_prices_product__variant_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.product_prices
    ADD CONSTRAINT product_prices_product__variant_id_fkey FOREIGN KEY (product__variant_id) REFERENCES public.product_variants(id);


--
-- Name: product_ratings product_ratings_product_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.product_ratings
    ADD CONSTRAINT product_ratings_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id);


--
-- Name: product_variant_colors product_variant_colors_color_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.product_variant_colors
    ADD CONSTRAINT product_variant_colors_color_id_fkey FOREIGN KEY (color_id) REFERENCES public.colors(id);


--
-- Name: product_variant_colors product_variant_colors_product_variant_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.product_variant_colors
    ADD CONSTRAINT product_variant_colors_product_variant_id_fkey FOREIGN KEY (product_variant_id) REFERENCES public.product_variants(id);


--
-- Name: product_variant_sizes product_variant_sizes_product_variant_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.product_variant_sizes
    ADD CONSTRAINT product_variant_sizes_product_variant_id_fkey FOREIGN KEY (product_variant_id) REFERENCES public.product_variants(id);


--
-- Name: product_variant_sizes product_variant_sizes_size_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.product_variant_sizes
    ADD CONSTRAINT product_variant_sizes_size_id_fkey FOREIGN KEY (size_id) REFERENCES public.sizes(id);


--
-- Name: product_variants product_variants_product_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.product_variants
    ADD CONSTRAINT product_variants_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id);


--
-- Name: products products_brand_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_brand_id_fkey FOREIGN KEY (brand_id) REFERENCES public.brands(id);


--
-- Name: products products_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.categories(id);


--
-- PostgreSQL database dump complete
--
