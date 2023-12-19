CREATE TABLE class_details (
  id SERIAL PRIMARY KEY,
  result VARCHAR(50) NOT NULL,
  description TEXT,
  symptom TEXT,
  cause TEXT,
  treatment TEXT, 
  prevention TEXT
);

INSERT INTO class_details (result, description, symptom, cause, treatment, prevention)
VALUES
(
  'Bacterial Diseases - Aeromoniasis',
  'Aeromoniasis adalah penyakit bakteri yang disebabkan oleh bakteri Aeromonas salmonicida. Penyakit ini sering menyerang ikan, terutama ikan koi dan ikan mas. Berikut adalah beberapa informasi penting tentang Aeromoniasis:', 
  'Infeksi bakteri Aeromonas mempengaruhi berbagai sistem dalam tubuh ikan, menghasilkan gejala seperti mata membesar (exophthalmos), penumpukan cairan di perut (ascitis), kerusakan ginjal (renal dropsy), dan sirip yang robek. Sebagian besar ikan yang terinfeksi memiliki kemerahan pada tubuh, dengan bercak perdarahan di insang, ekor, sirip, dinding tubuh, dan organ internal hewan, sementara yang lain menunjukkan ulkus kulit dan insang.',
  'Meskipun bakteri Aeromonas salmonicida menyebabkan infeksi, cedera, perubahan musiman, perubahan tajam dalam suhu air, dan sanitasi atau nutrisi yang buruk dapat membuat ikan lebih rentan terhadap bakteri.', 
  'Bergantung pada jenis bakteri Aeromonas yang dimiliki ikan, dokter hewan akan meresepkan obat untuk menghilangkan infeksi - biasanya antibiotik. Obat ini bisa disuntikkan ke dalam ikan atau ditambahkan ke air ikan.', 
  'Penting untuk diingat bahwa kondisi stres seperti kepadatan populasi, penanganan dan transportasi ikan, perubahan suhu yang signifikan, hipoksia, atau kondisi stres lainnya seringkali memicu wabah penyakit bakteri. Oleh karena itu, kontrol berbasis penghapusan faktor predisposisi sangat penting.'
),
(
  'Bacterial Gill Disease', 
  'Bacterial gill disease (BGD) adalah penyakit bakteri yang biasanya disebabkan oleh bakteri filamen dalam genus Flavobacterium, seringkali F.Branchiophilum. Penyakit ini biasanya menyerang ikan, terutama salmonid yang dibudidayakan. Berikut adalah beberapa informasi penting tentang Bacterial Gill Disease',
  'Ikan dengan BGD biasanya kehilangan nafsu makan, berorientasi pada arus air untuk meningkatkan aliran di insang, dan menunjukkan gerakan operkular yang berlebihan. Mungkin juga terjadi peningkatan lendir di kepala dan tubuh bagian atas. BGD biasanya mempengaruhi salmonid berukuran kecil atau fingerling dalam kondisi budidaya berdensitas tinggi.', 
  'Penyakit ini paling sering disebabkan oleh bakteri filamen dalam genus Flavobacterium. Faktor lingkungan seperti kepadatan populasi yang tinggi, kualitas air dan lingkungan yang sub-optimal seperti amonia berlebih, kadar oksigen terlarut rendah, dan materi organik tersuspensi berlebih dapat memicu penyakit ini.',
  'Intervensi dini dalam perkembangan penyakit dapat mengurangi kematian ikan yang bisa sangat signifikan. Dalam pengaturan hatchery, pengobatan kimia eksternal dengan hidrogen peroksida dapat membantu mengendalikan bakteri.',
  'Penting untuk diingat bahwa kondisi lingkungan yang buruk dan stres dapat memicu wabah penyakit ini. Oleh karena itu, kontrol berbasis penghapusan faktor predisposisi sangat penting.'
),
(
  'Bacterial Red Disease', 
  'Bacterial red disease adalah kondisi umum yang biasanya terjadi pada ikan air tawar, khususnya ikan permainan seperti bluegill (bream), largemouth bass, dan striped bass dan hibridanya. Berikut adalah beberapa informasi penting tentang Bacterial Red Disease:',
  'Kondisi ini ditandai dengan munculnya luka merah dan bisul pada ikan. Dalam bentuk yang paling ringan, kondisi ini terlihat sebagai “bisul” merah, atau lesi, pada ujung sirip, terutama sirip dorsal bluegill. Seiring perkembangan penyakit, ikan mungkin mengalami erosi sirip, dan bisul di sisi tubuh mereka.',
  'Bacterial Red Disease biasanya disebabkan oleh dua organisme, Aeromonas hydrophila, sebuah bakteri, dan Heteropolaria sp. (sebelumnya Epistylis sp.), sebuah protozoa. Kedua organisme ini ditemukan di lingkungan akuatik dan mampu menyebabkan penyakit.', 
  'Penyakit ini seringkali berjalan dan ikan mungkin pulih tanpa pengobatan. Namun, jika penyakit ini mencapai proporsi epidemi, berkontribusi pada kematian ikan yang signifikan (lebih dari 10 persen) dari ikan permainan, maka pengobatan diperlukan.',
  'Penting untuk diingat bahwa Bacterial Red Disease adalah istilah generik yang menggambarkan kondisi fisik ikan daripada merujuk pada agen penyakit spesifik. Oleh karena itu, ikan yang terkena harus dikirim ke laboratorium diagnostik penyakit ikan untuk mengidentifikasi dengan benar patogen yang berkontribusi pada setiap wabah penyakit.'
),
(
  'Fungal Diseases', 
  'Fungal diseases pada ikan adalah kondisi umum yang biasanya terjadi pada ikan air tawar. Penyakit ini disebabkan oleh spora jamur yang ada di lingkungan akuatik. Berikut adalah beberapa informasi penting tentang penyakit jamur pada ikan:',
  'Penyakit jamur biasanya mempengaruhi kulit dan insang ikan. Gejala umumnya termasuk pertumbuhan seperti kapas pada kulit atau insang, yang biasanya dimulai sebagai infeksi fokal kecil yang dapat menyebar dengan cepat ke seluruh tubuh.',
  'Spora jamur seperti Saprolegnia, Achlya, dan Fusarium adalah penyebab umum dari penyakit jamur pada ikan. Faktor lingkungan seperti kualitas air yang buruk dan stres dapat memicu penyakit ini.', 
  'Pengobatan untuk penyakit jamur pada ikan biasanya melibatkan penggunaan obat antijamur. Dalam kasus Saprolegniasis, misalnya, ikan dapat dibersihkan dengan larutan malachite green yang kuat dan kemudian diberi krim tahan air.',
  'Penting untuk diingat bahwa spora jamur berbahaya selalu ada di akuarium, tetapi biasanya tidak menimbulkan efek berbahaya kecuali ikan mengalami stres atau kondisi lingkungan memburuk.'
),
(
  'Healthy Fish', 
  'Ikan yang sehat adalah ikan yang memiliki penampilan dan perilaku yang normal. Ikan yang sehat memiliki warna yang cerah, sirip yang utuh, dan gerakan yang lincah. Ikan yang sehat juga memiliki nafsu makan yang baik dan tidak menunjukkan tanda-tanda penyakit. Untuk menjaga kesehatan ikan, penting untuk memberikan perawatan yang baik, termasuk menjaga kebersihan kolam atau akuarium, memberikan makanan yang bergizi, dan menjaga suhu dan kualitas air.',
  '',
  '', 
  '',
  ''
),
(
  'Parasitic Diseases', 
  'Parasitic diseases pada ikan adalah kondisi yang disebabkan oleh berbagai jenis parasit yang ada di lingkungan akuatik. Berikut adalah beberapa informasi penting tentang penyakit parasit pada ikan:',
  'Gejala penyakit parasit pada ikan bervariasi tergantung pada jenis parasit. Misalnya, parasit seperti Ichthyophthirius (air tawar) dan Cryptocaryon (air asin) menyebabkan bintik putih yang terlihat pada kulit atau sirip. Parasit lain seperti Trichodina (air tawar, air asin), Chilodonella (air tawar), dan Brooklynella (air asin) menyebabkan peningkatan laju pernapasan, pipa, kelebihan lendir, dan kehilangan kondisi.',
  'Semua kelompok utama parasit hewan ditemukan pada ikan, dan ikan liar yang tampak sehat sering kali membawa beban parasit yang berat. Parasit dengan siklus hidup langsung dapat menjadi patogen penting dari ikan budidaya; parasit dengan siklus hidup tidak langsung sering menggunakan ikan sebagai inang antara.', 
  'Pengobatan untuk penyakit parasit pada ikan biasanya melibatkan penggunaan obat antiparasit. Misalnya, dalam kasus infeksi Ichthyophthirius, ikan dapat dibersihkan dengan larutan malachite green yang kuat dan kemudian diberi krim tahan air.',
  'Penting untuk diingat bahwa parasit adalah bagian alami dari ekosistem akuatik dan biasanya tidak menimbulkan efek berbahaya kecuali ikan mengalami stres atau kondisi lingkungan memburuk.'
),
(
  'White Tail Disease', 
  'White tail disease (WTD) adalah penyakit ikan yang disebabkan oleh parasit Ichthyophthirius multifiliis (I. multifiliis). Parasit ini bersifat obligat parasitik, artinya virus ini hanya dapat hidup dan berkembang biak di dalam tubuh ikan. WTD menyerang berbagai jenis ikan air tawar, termasuk ikan lele, ikan nila, ikan mas, ikan gurame, dan ikan konsumsi lainnya. Penyakit ini dapat menyerang ikan di semua ukuran, tetapi ikan berukuran kecil lebih rentan terserang.  Berikut adalah beberapa informasi penting tentang penyakit WTD pada ikan:',
  'Gejala klinis WTD yang paling khas adalah munculnya bintik-bintik putih berukuran 1-2 mm pada kulit dan sirip ikan. Bintik-bintik ini sebenarnya adalah bentuk inang parasit I. multifiliis. Selain bintik-bintik putih, ikan yang terserang WTD juga akan menunjukkan gejala-gejala lain, seperti ikan menjadi lesu dan tidak aktif, ikan sering menggosok-gosokkan tubuhnya pada benda di sekitarnya, ikan bernafas dengan cepat, ikan kehilangan nafsu makan ,dan ikan mengalami luka-luka pada kulit dan sirip',
  'WTD dapat menyebar dengan cepat melalui kontak langsung antara ikan yang sakit dengan ikan yang sehat. Penyebaran penyakit ini juga dapat terjadi melalui media air, peralatan, dan manusia.', 
  'Obat-obatan yang digunakan untuk mengobati WTD adalah obat-obatan yang mengandung formalin, malachite green, atau copper sulphate. Obat-obatan ini dapat diberikan secara perendaman atau pencelupan.',
  'Untuk mencegah penyebaran WTD, dapat dilakukan beberapa upaya, antara lain menjaga kebersihan kolam atau akuarium, memantau kesehatan ikan secara rutin, melakukan karantina terhadap ikan baru yang akan dimasukkan ke kolam atau akuarium, menggunakan obat-obatan untuk mengobati ikan yang terserang WTD.'
);