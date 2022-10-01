<script lang="ts">
	import type { Article } from '../common/types';
	import NewsEntry from './news-entry.svelte';

	export let articles: Article[] = [];
</script>

<article class="col-md-12">
	<div class="cards-1 section-gray">
		<div class="container">
			<div class="row justify-content-center">
				{#each articles as article, index}
					<div class="col-md-5 col-sm-10 mb-3">
						<NewsEntry {article} articleModalId={`article-modal-${index}`} />
					</div>
				{/each}
			</div>
		</div>
	</div>
</article>

<style>
/*---------------------------------------------------------------------- /
SECTIONS
----------------------------------------------------------------------- */
	.section-gray {
		
		padding: 60px 0 30px 0;
	}
</style>
