import { Page } from "../components/page"
import { Link } from "react-router-dom";

export function Settings() {
    return (
        <Page>
            <div class="container text-center">
                <div class="col">
                    <Link to="/change-password" class="btn btn-primary">Change Password</Link>
                </div>
                <br></br>
                <div class="col">
                    <Link to="/change-email" class="btn btn-primary">Change Email</Link>
                </div>
            </div>
        </Page>
    )
}