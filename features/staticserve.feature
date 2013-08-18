Feature: staticserve command

  In order to provide help to the user
  As a command line user of staticserve
  I want to get help documentation on options

  Scenario: staticserve -help
    When I run `./staticserve -help`
    Then the exit status should be 2

  Scenario: staticserve -tls
    When I run `./staticserve -tls`
    Then the exit status should be 1

  Scenario: staticserve -tls -key KEYFILE
    When I run `./staticserve -tls -key KEYFILE`
    Then the exit status should be 1

  Scenario: staticserve -tls -cert CRTFILE
    When I run `./staticserve -tls -cert CRTFILE`
    Then the exit status should be 1
