import { Page } from "../components/page"

export function ForgotPassword() {
    return(
        <Page>
            <div class="container text-center w-25">
                <h3>Reset your password</h3>

                <form>
                    <div class="mb-3">
                        <label for="email-address" class="form-label">Email-address</label>
                        <input required type="text" placeholder="email" id="email" class="form-control"/>
                    </div>
                    <button type="submit" class="btn btn-primary">Submit</button>
                </form>

                <form>
                    <div class="mb-3">
                        <label for="email-address" class="form-label">Email-address</label>
                        <input required type="text" placeholder="email" id="email" class="form-control"/>
                    </div>
                    <button type="submit" class="btn btn-primary">Submit</button>
                </form>
            </div>
        </Page>
    )
}