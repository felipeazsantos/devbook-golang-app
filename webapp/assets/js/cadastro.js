$('#formulario-cadastro').on("submit", criarUsuario)

function criarUsuario(evento) {
    evento.preventDefault()
    console.log("Dentro da função usuário")

    if ($('#senha').val() != $('#confirmar-senha').val()) {
        alert("As senha não coincidem!")
        return
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: $("#nome").val(),
            email: $("#email").val(),
            nick: $("#nick").val(),
            senha: $("#senha").val(),
        }
    }).done(function() {
        alert("Usuáriao cadastrado com sucesso")
    }).fail(function () {
        alert("Erro ao cadastrar o usuário")
    });
}