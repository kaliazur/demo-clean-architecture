Feature: user creation
  In order to sell apis
  As a very powerful startup

  Scenario: Non admin User can Purchase
    Given non admin user with ID "42", username "kaliazur", firstname "Pierre", age 30
    Given 0 apis
    When I add 5 apis
    When user with ID "42" purchase 2 api
    Then there should be 5 remaining api

  Scenario: Non admin User can Purchase
    Given admin user with ID "42", username "kaliazur", firstname "Pierre", age 30
    Given 0 apis
    When I add 5 apis
    When user with ID "42" purchase 2 api
    Then there should be 3 remaining api
