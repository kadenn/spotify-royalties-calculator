import React from 'react';
import logo from './logo.svg';
import './App.css';
import { Playlists } from './components/playlists';
import { SpotifyUrlBar } from './components/spotifyUrlBar';
import { TrackResults } from './components/trackResults';
import 'bulma/css/bulma.min.css';

function App() {
  return (
    <div>

      <div className="header section">
        <div className="content is-large">
          <h1 className="title">Spotify Royalties Calculator</h1>
        </div>
        <img src={logo} className="app-logo" alt="logo" />
      </div>

      <div className="content container is-large is-fluid is-three-quarters-mobile">
        <p>
          Spotify Royalties Calculator is a simplified version of <strong>Blokur</strong> that uses the Spotify API to get song
          and artist data. The app basically estimates the total amount Spotify must pay the rights holders for a given song,
          album or playlist based on the data it receives from the Spotify API and calculates the amount for each right holder.
          For simplicity, it is assumed that all rights holders of the song will be given an equal share. Rights holders
          other than the artists shared by Spotify have been ignored.
        </p>
      </div>

      <div className="content container is-large">
        <Playlists />
        <br />
        <SpotifyUrlBar />
        <br />
        <TrackResults />
        <br />
      </div>

    </div >
  );
}

export default App;
