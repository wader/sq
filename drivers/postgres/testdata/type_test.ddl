create table type_test
(
    col_id                  serial primary key                                                                  not null,
    col_bigint              bigint                   default 0                                                  not null,
    col_bigint_n            bigint,
    col_bigserial           bigserial,
    col_bit                 bit                      default '0'::bit                                           not null,
    col_bit_n               bit,
    col_bitvarying          bit varying(255)          default '0'::bit                                           not null,
    col_bitvarying_n        bit varying(255),
    col_boolean             boolean                  default false                                              not null,
    col_boolean_n           boolean,
    col_box                 box                      default '(0,0), (0,0)'::box                                not null,
    col_box_n               box,
    col_bytea               bytea                    default ''::bytea                                          not null,
    col_bytea_n             bytea,
    col_character           character(255)           default ''                                                 not null,
    col_character_n         character(255),
    col_character_varying   character varying(255)   default ''                                                 not null,
    col_character_varying_n character varying(255),
    col_cidr                cidr                     default '0/0'::cidr                                        not null,
    col_cidr_n              cidr,
    col_circle              circle                   default '((0,0), 0)'::circle                               not null,
    col_circle_n            circle,
    col_date                date                     default '1970-01-01'::date                                 not null,
    col_date_n              date,
    col_double_precision    double precision         default 0                                                  not null,
    col_double_precision_n  double precision,
    col_inet                inet                     default '0/0'::inet                                        not null,
    col_inet_n              inet,
    col_integer             integer                  default 0                                                  not null,
    col_integer_n           integer,
    col_interval            interval                 default '0s'::interval                                     not null,
    col_interval_n          interval,
    col_json                json                     default '{}'::json                                         not null,
    col_json_n              json,
    col_jsonb               jsonb                    default '{}'::jsonb                                        not null,
    col_jsonb_n             jsonb,
    col_line                line                     default '[(0,0),(1,0)]'::line                              not null,
    col_line_n              line,
    col_lseg                lseg                     default '[(0,0),(0,0)]'::lseg                              not null,
    col_lseg_n              lseg,
    col_macaddr             macaddr                  default '00:00:00:00:00:00'                                not null,
    col_macaddr_n           macaddr,
    col_money               money                    default '0.0'::money                                       not null,
    col_money_n             money,
    col_numeric             numeric                  default 0.0                                                not null,
    col_numeric_n           numeric,
    col_path                path                     default '[(0,0)]'::path                                    not null,
    col_path_n              path,
    col_pg_lsn              pg_lsn                   default '0/0'::pg_lsn                                      not null,
    col_pg_lsn_n            pg_lsn,
    col_point               point                    default '(0,0)'::point                                     not null,
    col_point_n             point,
    col_polygon             polygon                  default '((0,0))'::polygon                                 not null,
    col_polygon_n           polygon,
    col_real                real                     default 0                                                  not null,
    col_real_n              real,
    col_smallint            smallint                 default 0                                                  not null,
    col_smallint_n          smallint,
    col_smallserial         smallserial, -- no col_smallserial_n, because serial cannot be NULL
    col_serial              serial,      -- no col_serial_n, because serial cannot be NULL
    col_text                text                     default ''                                                 not null,
    col_text_n              text,
    col_time                time                     default '00:00:00'::time without time zone                 not null,
    col_time_n              time,
    col_timetz              time with time zone      default '00:00:00+00'::time with time zone                 not null,
    col_timetz_n            time with time zone,
    col_timestamp           timestamp                default '1970-01-01 00:00:00'::timestamp without time zone not null,
    col_timestamp_n         timestamp,
    col_timestamptz         timestamp with time zone default '1970-01-01 00:00:00+00'::timestamp with time zone not null,
    col_timestamptz_n       timestamp with time zone,
    col_tsquery             tsquery                  default ''                                                 not null,
    col_tsquery_n           tsquery,
    col_tsvector            tsvector                 default ''                                                 not null,
    col_tsvector_n          tsvector,
    col_uuid                uuid                     default '00000000-0000-0000-0000-000000000000'::uuid       not null,
    col_uuid_n              uuid,
    col_xml                 xml                      default ''::xml                                            not null,
    col_xml_n               xml
);
