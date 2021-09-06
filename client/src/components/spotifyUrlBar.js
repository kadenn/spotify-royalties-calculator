import { useState } from "react";
import { useDispatch, useSelector } from 'react-redux';
import { appendTrack } from "../store/tracks";
import axios from "axios";

export function SpotifyUrlBar() {
    const [spotifyUrl, setSpotifyUrl] = useState("");
    const dispatch = useDispatch();
    const { data } = useSelector((store) => store.tracks);

    function checkAndAppendNewTrack(trackID) {
        var storedTrackIDs = data.map(function (track) { return track.trackID; });
        if (!storedTrackIDs.includes(trackID)) {
            dispatch(appendTrack(trackID));
        }
    }

    const submitSpotifyUrl = () => {
        try {
            var urlType = spotifyUrl.split('/')[3];
            var spotifyID = spotifyUrl.split('/')[4].split('?')[0];

            if (urlType === 'track') {
                checkAndAppendNewTrack(spotifyID);

            } else if (urlType === 'playlist') {
                axios.get(`${process.env.REACT_APP_API_URL}/playlist_tracks/${spotifyID}`).then(resp => {
                    resp.data.tracks.forEach(track => {
                        checkAndAppendNewTrack(track.trackID);
                    });
                });

            } else if (urlType === 'album') {
                axios.get(`${process.env.REACT_APP_API_URL}/album_tracks/${spotifyID}`).then(resp => {
                    resp.data.tracks.forEach(track => {
                        checkAndAppendNewTrack(track.trackID);
                    });
                });
            }

        }
        catch (err) {
            alert("Please enter a valid Spotify URL!");
        }
        setTimeout(function () {
            document.getElementById("results").scrollIntoView();
        }, 1000);
    }

    return (
        <div className="box">
            <div className="content container is-large">
                <p><strong>You can also enter the Spotify URL of a song, album or playlist</strong></p>
                <input className="input is-large" type="text" onChange={e => setSpotifyUrl(e.target.value)} placeholder="https://open.spotify.com/track/4TXYAETrGC53xgXZ7ykNEl?si=7e04081882ee4d8b" />
            </div>
            <button className="button is-info is-large" onClick={submitSpotifyUrl}>Calculate</button>
        </div>
    );
};