<script>
	import Card from '$lib/components/card.svelte';
	import Request from '$lib/components/request.svelte';
	export let data;

	let requests = data.requests,
		error = data.error;
</script>

<Card class="w-75">
	<h1 class="text-center">Solicitações de empréstimo de livros</h1>

	<div class="d-flex justify-content-center mt-5 px-5">
		<table class="table table-striped">
			<thead>
				<tr>
					<th>Solicitações</th>
				</tr>
			</thead>
			<tbody>
				{#if requests.length}
					{#each requests as request}
						<tr>
							<td>
								<Request {request} />
							</td>
						</tr>
					{/each}
				{:else}
					<tr>
						<td>Não há solicitações em aberto!</td>
					</tr>
				{/if}
			</tbody>
		</table>
	</div>
</Card>

<style>
</style>
