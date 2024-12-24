# Chapter 5: Structuring Command-Line Programs

This chapter explores two key approaches to flag parsing:
  - Using the `os` package to handle raw command-line input.
  - Using the `flag` package for structured flag parsing.

We'll first build a custom parser using the `os` package to gain a deeper understanding of how flag parsing works under the hood and help us explore more advanced Go concepts, such as escape analysis, memory management, error handling, and so on.

**Topics covered:**
- Crafting idiomatic and user-friendly command-line programs.
- Parsing command-line arguments and flags.
- Extending the flag package's functionality by satisfying interfaces.
- Retrieving and using positional arguments for mandatory arguments.
- Validating command-line arguments and flags.

## Listings

- Laying out the groundwork
  - [Listing 5.1: Printing the logo and usage message](../../all-listings/05-structuring-command-line-programs/01-printing-the-logo-and-usage-message.md)
- Crafting a custom flag parser
  - [Listing 5.2: Implementing the config type](../../all-listings/05-structuring-command-line-programs/02-implementing-the-config-type.md)
  - [Listing 5.3: Implementing the flag value parsers](../../all-listings/05-structuring-command-line-programs/03-implementing-the-flag-value-parsers.md)
  - [Listing 5.4: Implementing the flag parser](../../all-listings/05-structuring-command-line-programs/04-implementing-the-flag-parser.md)
  - [Listing 5.5: Integrating the custom parser](../../all-listings/05-structuring-command-line-programs/05-integrating-the-custom-parser.md)
  - [Listing 5.6: Setting sensible defaults](../../all-listings/05-structuring-command-line-programs/06-setting-sensible-defaults.md)
- The flag package
  - [Listing 5.7: Parsing flags using a flag set](../../all-listings/05-structuring-command-line-programs/07-parsing-flags-using-a-flag-set.md)
  - [Listing 5.8: Removing error and usage messages](../../all-listings/05-structuring-command-line-programs/08-removing-error-and-usage-messages.md)
- Custom flag types
  - [Listing 5.9: Implementing a custom dynamic value](../../all-listings/05-structuring-command-line-programs/09-implementing-a-custom-dynamic-value.md)
  - [Listing 5.10: Defining a custom flag type](../../all-listings/05-structuring-command-line-programs/10-defining-a-custom-flag-type.md)
- Positional arguments
  - [Listing 5.11: Switching to a positional argument](../../all-listings/05-structuring-command-line-programs/11-switching-to-a-positional-argument.md)
- Post-parse flag validation
  - [Listing 5.12: Implementing a custom validator](../../all-listings/05-structuring-command-line-programs/12-implementing-a-custom-validator.md)
  - [Listing 5.13: Integrating the custom validator](../../all-listings/05-structuring-command-line-programs/13-integrating-the-custom-validator.md)
