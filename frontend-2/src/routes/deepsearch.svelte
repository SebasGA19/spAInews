<script>
import Page from '../components/page.svelte';

</script>

<Page>
    <div class="container text-center" style="max-width: 50vw;">
        <form>
            <div class="form-group">
              <label for="exampleInputEmail1">Text Search</label>
              <input type="text" class="form-control" id="exampleInputEmail1" aria-describedby="emailHelp" placeholder="Enter email">
              
            </div>
            <div class="form-group">
              <label for="exampleInputPassword1">Title/Key words search</label>
              <input type="text" class="form-control" id="exampleInputPassword1" placeholder="Password">
            </div>
            First last news
            <div class="form-check form-switch">
                <label class="form-check-label">
                    <input class="form-check-input" type="checkbox" role="switch">
                </label>
            </div>
            <button type="submit" class="btn btn-primary">Submit</button>

          </form>
          
    </div>
</Page>