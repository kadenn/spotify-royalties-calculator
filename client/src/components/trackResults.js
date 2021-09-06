import { useSelector } from 'react-redux';

export function TrackResults() {
    const { data } = useSelector((store) => store.tracks);

    return (
        <div id="results">
            {data?.length > 0 &&
                <form className="box">
                    <div className="content container is-large">
                        <p><strong>Results</strong></p>
                    </div>

                    <div className="columns is-multiline">
                        {data.map(track =>
                            <div className="column is-12" key={track.trackID}>
                                <div className="box">
                                    <article className="media">
                                        <div className="media-left image is-128x128">
                                            <img src={track.images[0]?.url} alt="track" />
                                        </div>
                                        <div className="media-content">
                                            <strong>{track.name}</strong>
                                            <br />
                                            The song was played <strong>{track.playCount}</strong> times.
                                            <br />
                                            A total of <strong>£{(track.totalRoyalties).toFixed(2)}</strong> royalties were paid.
                                            <br />
                                            {track.artists.map(artist =>
                                                <div className="box" key={artist.id}>
                                                    <strong>{artist.name}</strong>
                                                    <br />
                                                    <progress value={parseInt(100 / track.artists.length)} max="100">{parseInt(100 / track.artists.length)}</progress> <strong>{parseInt(100 / track.artists.length)}%</strong>
                                                    <br />
                                                    {artist.name} earned <strong>£{(track.totalRoyalties / track.artists.length).toFixed(2)}</strong> from this song.
                                                </div>
                                            )}
                                        </div>
                                    </article>
                                </div>

                            </div>
                        )}

                    </div>
                </form>
            }
        </div>
    );
};