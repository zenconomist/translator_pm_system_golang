# translator_pm_system_golang

Author's Note: this is an unfinished project.

    The current codebase is mainly for my codewriting demonstrating purpose, and
    the main features are coming from a former project, which is refactored and
    rewritten here (partially, by now).

------------------------------
The Project
    
    The Project was to create a web-based project management or CRM-system
    for a translation company.

    The Business:
        The business is for the company that they receive a relatively
        large amount of translation tasks, which vary from very small
        to very big projects, which are then distributed to outsourced
        and inbound translators, and several other tasks have to
        be managed by the company, like lektoring, reviewing, preparing
        the tasks, etc.

    The Original Solution:
        The original solution was a delphi-based desktop application 
        which had several restrictions.

------------------------------
Features

    Features of the new pm system:
        - Projects can have zero, one, or more tasks
        - the tasks can be assigned to suppliers / translators
        - translators and suppliers have their own access to the
            system, they can perform actions, like accepting 
            or declining the tasks, setting them as ready, etc.
        - the system is integrated to a 3rd party invoicing 
            supplier called Billingo (not implemented here)
        - the system is integrated to the company's sharepoint
            through MS Graph API. (implemented)
            (Note: The original code in the msgraph package 
            was written by someone else, which included estab-
            lishing msgraph client, and making calls. Small
            adjustments were made by the author.)
        - e-mails should be sent from the system based on
            certain conditions, to several stakeholders
        - several states of tasks, projects, timestates,
            prepare and review states are to be persisted,
            validated and statechanges can trigger actions.

------------------------------
Design concepts

    1. Layered Architecture
        The entrypoint of the system as a REST API
        is in the server package.
        The server package has GEN_ type files, which
        are generated with the codegenerator package,
        and can have custom handlers.
        The handlers interact with the received url
        vars or json values, extract them into dto-s
        and establish services.
        Services are generic services written in the
        services package and folder
        "A__services_layer00_Service.go" file is the
        generic service containing file.
        When generated type, custom mapping and validations
        on data happens first at the DTO level, when
        mapping a DTO into an Entity.
        At last, Services call Repositories, which are also
        mainly generic.

        Aside from generic and generated solutions,
        every layer can have custom solutions implemented.
        Custom handlers, with custom services, with custom
        repositories.

    2. SOLID principles
        Not strictly kept in mind, but trying :).
        Loose coupling and dependency inversion are kept
        in mind when writing interfaces for environment and
        logging packages, and using these interfaces in other
        packages, and keeping the main package to tell which
        exact type to use for the interfaces.

    3. Generics
        DRY principle -> instead of generating lots of code
        for the repository and the service layer, by most 
        entities the steps are the same when creating, 
        updating, deleting or retreiving them. For these
        purposes, generics-based services and generics-
        based repositories are perfect.

    4. Design patterns
        Trying to use some design patterns, like:
        - Factory
        - Builder
        - State Machine
            -> at State Machine the initial concept
            was challenged and changed to a more
            database-heavy approach in order to 
            gain higher flexibility. (This means,
            I wanted to keep already created tasks
            in the state-change pattern where they
            started, and only newly created tasks
            and projects have to adjust to the new
            state-change pattern, when it is changed.)

    5. Historicization
        Though not a basic feature in normal applications sometimes,
        historicization and keeping track of what happened in the
        database is a key feature in my opinion, so when applicable,
        Create - Update and Delete actions are historicized in an
        SCD2 fashion. 

    6. Testing
        Keeping in mind that test-writing will help
        maintain the code, so write as many unit-tests as
        possible (in the current phase, testing helped mainly
        the development phase in trying to mock certain 
        functions and features).