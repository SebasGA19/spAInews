import { Outlet, Link } from "react-router-dom";

export function Layout() {
    return (
        <>
            <nav class="navbar navbar-expand-lg bg-light">
                <div class="container-fluid">
                    <div class="collapse navbar-collapse" id="navbarNav">
                        <ul class="navbar-nav">
                            <li class="nav-item">
                                <Link to="/" class="nav-link">Home</Link>
                            </li>
                            <li class="nav-item">
                                <Link to="/login" class="nav-link">Login</Link>
                            </li>
                            <li class="nav-item">
                                <Link to="/register" class="nav-link">Register</Link>
                            </li>
                            <li class="nav-item">
                                <Link to="/confirm-register" class="nav-link">Confirm registration</Link>
                            </li>
                        </ul>
                    </div>
                </div>
            </nav>
            <Outlet />
        </>
    )
};
