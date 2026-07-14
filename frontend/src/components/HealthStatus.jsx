import { useEffect, useState } from "react";
import { checkHealth } from "../api";

function HealthStatus() {
    const [online, setOnline] = useState(false);

    useEffect(() => {
        async function load() {
            try {
                await checkHealth();
                setOnline(true);
            } catch {
                setOnline(false);
            }
        }

        load();
    }, []);

    return (
        <section className="card">
            <h2>API Status</h2>

            {online ? (
                <p className="online">🟢 API Online</p>
            ) : (
                <p className="offline">🔴 API Offline</p>
            )}
        </section>
    );
}

export default HealthStatus;