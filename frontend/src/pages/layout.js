import { Outlet, Link } from "react-router-dom";

export function Layout() {
    return (
        <>
            <nav class="navbar navbar-expand-lg bg-light">
                <div class="container-fluid">
                    <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
                        <span class="navbar-toggler-icon"></span>
                    </button>
                    <div class="collapse navbar-collapse" id="navbarNav">
                        <ul class="navbar-nav">
                            <li class="nav-item">
                                <Link to="/" class="nav-link">spAInews</Link>
                            </li>
                            <li class="nav-item">
                                <Link to="/login" class="nav-link">Login</Link>
                            </li>
                            <li class="nav-item">
                                <Link to="/register" class="nav-link">Register</Link>
                            </li>
                            <li class="nav-item">
                                <Link to="/settings" class="nav-link">Settings</Link>
                            </li>
                        </ul>
                    </div>
                </div>
            </nav>
            <Outlet />
        </>
    )
};
