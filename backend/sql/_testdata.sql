INSERT INTO repos (
    id,
    name,
    description,
    url,
    homepage
) VALUES (
    'MDEwOlJlcG9zaXRvcnkzMDY1NDU0',
    'impress.js',
    'Its a presentation framework',
    'https://github.com/impress/impress.js',
    'http://impress.js.org'
), (
    'MDEwOlJlcG9zaXRvcnkyNDI4MTI1ODc=',
    'shox',
    'A customisable, universally compatible terminal status bar',
    'https://github.com/liamg/shox',
    ''
), (
    'MDEwOlJlcG9zaXRvcnkxMTY5NTE2NTg=',
    'Astra',
    'Automated Security Testing For REST APIs',
    'https://github.com/flipkart-incubator/Astra',
    ''
);

INSERT INTO tags (id, name) VALUES (1, 'golang'), (2, 'python'), (3, 'rust');

INSERT INTO mapping (repo_id, tag_id) VALUES
('MDEwOlJlcG9zaXRvcnkzMDY1NDU0', 1),
('MDEwOlJlcG9zaXRvcnkzMDY1NDU0', 3),
('MDEwOlJlcG9zaXRvcnkyNDI4MTI1ODc=', 2),
('MDEwOlJlcG9zaXRvcnkyNDI4MTI1ODc=', 3);
