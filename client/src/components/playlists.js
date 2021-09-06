import { useEffect } from "react";
import { useSelector, useDispatch } from 'react-redux';
import { loadPlaylists } from "../store/playlists";
import { appendTrack } from "../store/tracks";
import axios from "axios";

const useFetching = () => {
    const dispatch = useDispatch();
    useEffect(() => {
        dispatch(loadPlaylists())
    }, [dispatch])
}

export function Playlists() {
    useFetching();
    const { playlists } = useSelector((store) => store.playlists.data);
    const { data } = useSelector((store) => store.tracks);

    const dispatch = useDispatch();

    function checkAndAppendNewTrack(trackID) {
        var storedTrackIDs = data.map(function (track) { return track.trackID; });
        if (!storedTrackIDs.includes(trackID)) {
            dispatch(appendTrack(trackID));
        }
    }

    const submitPlaylistUrl = (playlistUrl) => {
        try {
            var playlistID = playlistUrl.split('/')[5];
            axios.get(`${process.env.REACT_APP_API_URL}/playlist_tracks/${playlistID}`).then(resp => {
                resp.data.tracks.forEach(track => {
                    checkAndAppendNewTrack(track.trackID);
                });
            });
        }
        catch (err) {
            alert("Something went wrong!");
        }

        setTimeout(function () {
            document.getElementById("results").scrollIntoView();
        }, 1000);

    }

    return (
        <div >
            {playlists?.length > 0 &&
                <div className="box">
                    <div className="content container is-large">
                        <p><strong>Click one of the featured playlists below to start</strong></p>
                    </div>

                    <div className="columns is-multiline is-mobile">
                        {playlists.map(playlist =>
                            <div className="column is-2" key={playlist.playlistID}>
                                <img className="is-clickable" onClick={() => submitPlaylistUrl(playlist.tracks.href)} src={playlist.images[0]?.url} alt="playlist" />
                            </div>)}
                    </div>
                </div>
            }
        </div>
    );
};