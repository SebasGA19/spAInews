<script lang="ts">
	import type { Article } from '../common/types';
	import NewsEntry from './news-entry.svelte';

	export let articles: Article[] = [];
</script>

<div class="container">
	<div class="row justify-content-center mt-5 mb-5">
		{#each articles as article, index}
			<div class="col-md-5 col-sm-10 mb-3">
				<NewsEntry article={article} articleModalId={`article-modal-${index}`}/>
			</div>
		{/each}
	</div>
</div>
