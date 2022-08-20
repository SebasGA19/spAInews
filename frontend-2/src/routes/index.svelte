<script>
import Page from "../components/page.svelte";

</script>

<Page> 
    <header class="header">
        <h3 id="titulo">Welcome to the Jungle</h3>
    </header>
        
       <main class="main columns">
        <section class="column main-column">
          <a class="article first-article" href="#">
            <figure class="article-image is-4by3">
              <img src="https://www.eltiempo.com/files/article_main/uploads/2022/08/19/63005011cda23.jpeg" alt="">
            </figure>
            <div class="article-body">
              <h2 class="article-title">
                Noticia 1
              </h2>
              <p class="article-content">
                Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis.
                Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros.
                <a class="article-news" href="#">Continuar leyendo</a>
              </p>
              <footer class="article-info">
                <span>By Joe Smith</span>
                <span class="date-published" >20/8/22</span>
              </footer>
            </div>
          </a>
      
          <div class="columns">
            <div class="column nested-column">
              <a class="article" href="#">
                <figure class="article-image is-16by9">
                  <img src="https://static01.nyt.com/images/2022/08/19/world/19norway-walrus1-esp-1/merlin_211448394_20267413-9a7a-4307-afa2-b5d451b1bffd-superJumbo.jpg?quality=75&auto=webp" alt="">
                </figure>
                <div class="article-body">
                  <h2 class="article-title">
                    Hello World
                  </h2>
                  <p class="article-content">
                    Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis.
                    Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis.
                    <a class="article-news"  href="#">Continuar leyendo</a>
                  </p>
                  <footer class="article-info">
                    <span>By Joe Smith</span>
                    <span class="date-published" >20/8/22</span>
                  </footer>
                </div>
              </a>
            </div>
      
            <div class="column">
              <a class="article" href="#">
                <figure class="article-image is-16by9">
                  <img src="https://www.eltiempo.com/files/article_main/uploads/2022/08/19/63003ec4980bb.jpeg" alt="">
                </figure>
                <div class="article-body">
                  <h2 class="article-title">
                    Hello World
                  </h2>
                  <p class="article-content">
                    Lorem ipsum dolor sit amet, consectetur adipiscing elit.
                    <a class="article-news"  href="#">Continuar leyendo</a>
                  </p>
                  <footer class="article-info">
                    <span>By Joe Smith</span>
                    <span class="date-published" >20/8/22</span>
                  </footer>
                </div>
              </a>
              <a class="article" href="#">
                <div class="article-body">
                  <h2 class="article-title">
                    Hello World
                  </h2>
                  <p class="article-content">
                    Lorem ipsum dolor sit amet.
                    <a class="article-news"  href="#">Continuar leyendo</a>
                  </p>
                  <footer class="article-info">
                    <span>By Joe Smith</span>
                    <span class="date-published" >20/8/22</span>
                  </footer>
                </div>
              </a>
              <a class="article" href="#">
                <div class="article-body">
                  <h2 class="article-title">
                    Hello World
                  </h2>
                  <p class="article-content">
                    Lorem ipsum dolor sit amet feugiat facilisis.
                    <a class="article-news" href="#">Continuar leyendo</a>
                  </p>
                  <footer class="article-info">
                    <span>By Joe Smith</span>
                    <span class="date-published" >20/8/22</span>
                  </footer>
                </div>
              </a>
            </div>
          </div>
        </section>
      
        <section class="column">
          <a class="article" href="#">
            <figure class="article-image is-3by2">
              <img src="https://s3-us-west-2.amazonaws.com/s.cdpn.io/210284/image-04-lo.jpg" alt="">
            </figure>
            <div class="article-body">
              <h2 class="article-title">
                Hello World
              </h2>
              <p class="article-content">
                Lorem ipsum dolor sit amet, consectetur adipiscing elit.
                Lorem ipsum dolor sit amet, consectetur adipiscing elit.
                Lorem ipsum dolor sit amet, consectetur adipiscing elit.
                <a class="article-news"  href="#">Continuar leyendo</a>
              </p>
              <footer class="article-info">
                <span>By Joe Smith</span>
                <span class="date-published" >20/8/22</span>
              </footer>
            </div>
          </a>
          <a class="article" href="#">
            <div class="article-body">
              <h2 class="article-title">
                Hello World
              </h2>
              <p class="article-content">
                Lorem ipsum dolor sit amet, consectetur adipiscing elit.
                <a class="article-news"  href="#">Continuar leyendo</a>
              </p>
              <footer class="article-info">
                <span>By Joe Smith</span>
                <span class="date-published" >20/8/22</span>
              </footer>
            </div>
          </a>
          <a class="article" href="#">
            <div class="article-body">
              <h2 class="article-title">
                Hello World
              </h2>
              <p class="article-content">
                Lorem ipsum dolor sit amet, consectetur adipiscing elit.
                <a class="article-news"  href="#">Continuar leyendo</a>
              </p>
              <footer class="article-info">
                <span>By Joe Smith</span>
                <span class="date-published" >20/8/22</span>
              </footer>
            </div>
          </a>
          <a class="article" href="#">
            <div class="article-body">
              <h2 class="article-title">
                Hello World
              </h2>
              <p class="article-content">
                Lorem ipsum dolor sit amet, consectetur adipiscing elit.
                <a class="article-news"  href="#">Continuar leyendo</a>
              </p>
              <footer class="article-info">
                <span>By Joe Smith</span>
                <span class="date-published" >20/8/22</span>
              </footer>
            </div>
          </a>
        </section>
      </main>
</Page>
<style>
    html {
  background: mediumseagreen;
  font-size: 14px;
  -moz-osx-font-smoothing: grayscale;
  -webkit-font-smoothing: antialiased;
  min-width: 300px;
  overflow-x: hidden;
  overflow-y: scroll;
  text-rendering: optimizeLegibility;
}

body {
  color: #666;
  font-family: "Source Sans Pro", "Helvetica", "Arial", sans-serif;
  font-size: 1rem;
  font-weight: 400;
  line-height: 1.4;
}

a {
  text-decoration: none;
  transition: all 0.3s cubic-bezier(.25, .8, .25, 1);
  

}

div, h2, p, figure {
  margin: 0;
  padding: 0;
}

.header {
  color: white;
  padding: 40px 0 20px;
  text-align: center;
}

.header h1 {
  font-size: 40px;
  font-weight: bold;
}

.header h2 a {
  border-bottom: 1px solid rgba(255, 255, 255, 0.5);
  color: white;
  font-size: 20px;
  opacity: 0.5;
}

.header h2 a:hover {
  border-bottom-color: white;
  opacity: 1;
}

.main {
  margin: 0 auto;
  max-width: 1200px;
  padding: 10px;
}

.column {
  flex: 1;
  flex-direction: column;
}

.article {
  background: white;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.12), 0 1px 2px rgba(0, 0, 0, 0.24);
  color: #666;
  display: flex;
  flex: 1;
  flex-direction: column;
  flex-basis: auto;
  margin: 10px;

}

.article:hover,
.article:focus {
  box-shadow: 0 14px 28px rgba(0, 0, 0, 0.25), 0 10px 10px rgba(0, 0, 0, 0.22);
  color: #444;
}

.article-image {
  background: #eee;
  display: block;
  padding-top: 75%;
  position: relative;
  width: 100%;
  height: 100%;
}

.article-image img {
  display: block;
  height: 100%;
  left: 0;
  position: absolute;
  top: 0;
  width: 100%;
}


.article-image.is-3by2 {
  padding-top: 66.6666%;
}

.article-image.is-16by9 {
  padding-top: 56.25%;
}

.article-body {
    display: flex;
    flex: 1;
    flex-direction: column;
    padding: 20px;
}   

.article-title {
  
  flex-shrink: 0;
  font-size: 1.4em;
  font-weight: 700;
  line-height: 1.2;
}

.article-content {
  flex: 1;
  margin-top: 5px;
}

.article-info {
  display: flex;
  font-size: 0.85em;
  justify-content: space-between;
  margin-top: 10px;
}
.article-news{
    background: linear-gradient(black, white);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    font-size: 18px;
}
.article-news :hover{
    border-bottom: 1px solid;
    background: gray;
}


@media screen and (min-width: 800px) {
  .columns,
  .column {
    display: flex;
  }
}
@media screen and (min-width: 1000px) {
  .first-article {
    flex-direction: row;
  }

  .first-article .article-body {
    flex: 1;
  }

  .first-article .article-image {
    height: 300px;
    order: 2;
    padding-top: 0;
    width: 400px;
  }

  .main-column {
    flex: 3;
  }

  .nested-column {
    flex: 2;
  }
}
</style>
