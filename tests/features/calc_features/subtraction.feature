Feature: Subtraction

    Scenario:
        Given calculator is cleared

        When i press 2
            And i add 5
            And i subtract 3
            And i multiply by 0

        Then the result should be 4