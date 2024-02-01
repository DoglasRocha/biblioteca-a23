<script>
	import Card from '$lib/components/card.svelte';
	import { blur } from 'svelte/transition';

	let hasLoan = true;
	let canRenew = true;
	let returnDate = new Date();
	let newReturnDate = new Date();
	let hasRenewed = false;

	newReturnDate.setDate(newReturnDate.getDate() + 7);
</script>

<Card class="w-75">
	<h1 class="text-center">Renovar Empréstimo</h1>
	<div>
		{#if hasLoan}
			{#if canRenew}
				<div class="mt-3" transition:blur>
					<p>
						Você pode renovar o livro <span class="fw-bold">{'João e o Pé de Feijão'}</span> até o
						dia
						<span class="fw-bold">{newReturnDate.toLocaleDateString('pt-BR')}</span>.
					</p>
				</div>
				<div class="d-flex justify-content-end">
					<button
						class="btn btn-dark"
						on:click={() => {
							hasRenewed = true;
							canRenew = false;
						}}>Renovar</button
					>
				</div>
			{:else if hasRenewed}
				<div transition:blur>
					<p>
						O livro <span class="fw-bold">{'João e o Pé de Feijão'}</span> foi renovado com sucesso!
						Nova data de devolução:
						<span class="fw-bold">{newReturnDate.toLocaleDateString('pt-BR')}</span>
					</p>
				</div>
			{:else}
				<div transition:blur>
					<p>
						Você não pode renovar o livro <span class="fw-bold">{'João e o Pé de Feijão'}</span>,
						você deve devolvê-lo dia
						<span class="fw-bold">{returnDate.toLocaleDateString('pt-BR')}</span>.
					</p>
				</div>
			{/if}
		{:else}
			<p>Você não possui empréstimo em aberto, portanto, não tem o que renovar.</p>
		{/if}
	</div>
</Card>
