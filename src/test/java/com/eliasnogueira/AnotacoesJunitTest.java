package com.eliasnogueira;

import org.junit.jupiter.api.*;

@TestMethodOrder(
       // MethodOrderer.MethodName.class
       // MethodOrderer.OrderAnnotation.class
        MethodOrderer.DisplayName.class
)
public class AnotacoesJunitTest {

    @BeforeAll
    static void precondicapPorTodaClasseTeste() {
        System.out.println("Precondição para todos os testes");
    }

    @BeforeEach
    void precondicapPorTeste() {
        System.out.println("Precondição para cada um dos testes");
    }

    @AfterAll
    static void poscondicaoPorTodaClasseTeste() {
        System.out.println("Poscondição para todos os testes");
    }

    @AfterEach
    void poscondicaoPorTeste() {
        System.out.println("Poscondição para cada um dos testes");
    }

    @Test
    //@Order(3)
    @DisplayName("Aqui vai meu teste A")
    void aMeuTeste() {
        System.out.println("Aqui vai meu teste A");
    }


    @Test
    //@Order(2)
    @DisplayName("Aqui vai meu teste B")
    void bMeuTeste() {
        System.out.println("Aqui vai meu teste B");
    }

    @Test
    //@Order(1)
    @DisplayName("Aqui vai meu teste C")
    void cMeuTeste() {
        System.out.println("Aqui vai meu teste C");
    }

}
