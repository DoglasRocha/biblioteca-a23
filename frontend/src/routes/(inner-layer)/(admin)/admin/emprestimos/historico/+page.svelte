<script>
	import Card from '$lib/components/card.svelte';
	import Loan from '../../../../../../lib/components/loan.svelte';

	export let data;
	let loans = data.loans,
		error = data.error;
</script>

<Card class="w-75">
	<h1 class="text-center">Histórico de empréstimo de livros</h1>

	<div class="d-flex justify-content-center mt-5 px-5">
		<table class="table table-striped">
			<thead>
				<th>Empréstimos</th>
			</thead>
			<tbody>
				{#if loans.length}
					{#each loans as loan}
						<tr>
							<td>
								<Loan {loan} />
							</td>
						</tr>
					{/each}
				{:else}
					<tr><td>Não há empréstimos.</td></tr>
				{/if}
				{#if error}
					<tr>
						<td class="text-danger">{error}</td>
					</tr>
				{/if}
			</tbody>
		</table>
	</div>
</Card>

<style>
</style>
