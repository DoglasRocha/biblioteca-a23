<script>
	import Card from '$lib/components/card.svelte';
	import BigLogo from '$lib/components/big-logo.svelte';
	import Link from '$lib/components/link.svelte';
</script>

<Card class="w-75">
	<h1 class="text-center">Administração</h1>
	<BigLogo />
	<p class="mt-3">
		Esta é a parte dos administradores. Aqui, vocês podem controlar toda a operação da biblioteca.
	</p>
	<p>
		Para cadastrar um livro, clique em <Link href="/admin/cadastrar">Cadastrar livro</Link> na barra
		de navegação.
	</p>
	<p>
		Para buscar livros cadastrados, editá-los ou apagá-los, clique em <Link href="/admin/buscar"
			>Buscar livro</Link
		> na barra de navegação.
	</p>
	<p>
		Para verificar quem são os outros administradores da biblioteca ou aprovar novos
		administradores, clique em <Link href="/admin/admins">Administradores</Link> na barra de navegação.
	</p>
	<p>
		Para verificar solicitações de empréstimo, empréstimos ativos ou o histórico de empréstimos,
		clique em <Link href="#top">Empréstimos</Link> e em seguida na opção subsequente desejada, na barra
		de navegação.
	</p>
	<p>
		Lorem ipsum dolor sit amet consectetur adipisicing elit. Nemo ab repellat unde vel provident,
		deleniti autem non qui, quos voluptatibus eveniet quaerat dolorem in tempore sed, minima
		adipisci totam recusandae!
	</p>
	<p>
		Lorem, ipsum dolor sit amet consectetur adipisicing elit. Sunt, consequuntur tempora. Obcaecati,
		dolore cum veritatis qui facere veniam! Aspernatur, error! Officiis ex dicta aliquam ipsa
		voluptatibus nostrum odit consectetur magni!
	</p>
</Card>
