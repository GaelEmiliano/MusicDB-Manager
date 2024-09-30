# MusicDBManager

**MusicDBManager** is a **music manager for songs in mp3 format**, using an
**SQLite** database to perform queries based on the **ID3v2** tags of the
music files. The sytem scans the '~/Music' directory, reads the **ID3v2** 
tags from mp3 files, and saves information such as the **title**, **artist**,
**album**, **genre**, **track number**, **year**, and **album art** to the
database..

The system offers a graphical interface (GUI) for viewing and modifying the
database entries, without altering the origin tags of the music files.
Additionally, a **simple query language** is provided to perform advanced
searches through the GUI.

## Technologies
- **Language**: Go '>=1.20'
- **Database**: SQLite
