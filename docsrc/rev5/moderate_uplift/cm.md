# CM - Configuration Management

* Controls Count: 15
* Controls IDs: CM-12, CM-12 (1), CM-2 (2), CM-2 (3), CM-2 (7), CM-3, CM-3 (2), CM-3 (4), CM-4 (2), CM-7 (1), CM-7 (2), CM-7 (5), CM-8 (1), CM-8 (3), CM-9

## Controls

### CM-2 (2): Automation Support for Accuracy and Currency

Maintain the currency, completeness, accuracy, and availability of the baseline configuration of the system using automated mechanisms for maintaining baseline configuration of the system are defined;.

Automated mechanisms that help organizations maintain consistent baseline configurations for systems include configuration management tools, hardware, software, firmware inventory tools, and network management tools. Automated tools can be used at the organization level, mission and business process level, or system level on workstations, servers, notebook computers, network components, or mobile devices. Tools can be used to track version numbers on operating systems, applications, types of software installed, and current patch levels. Automation support for accuracy and currency can be satisfied by the implementation of [CM-8(2)](#cm-8.2) for organizations that combine system component inventory and baseline configuration activities.

the currency of the baseline configuration of the system is maintained using automated mechanisms for maintaining baseline configuration of the system are defined;;

the completeness of the baseline configuration of the system is maintained using automated mechanisms for maintaining baseline configuration of the system are defined;;

the accuracy of the baseline configuration of the system is maintained using automated mechanisms for maintaining baseline configuration of the system are defined;;

the availability of the baseline configuration of the system is maintained using automated mechanisms for maintaining baseline configuration of the system are defined;.

Configuration management policy

procedures addressing the baseline configuration of the system

configuration management plan

system design documentation

system architecture and configuration documentation

system configuration settings and associated documentation

system component inventory

configuration change control records

system security plan

other relevant documents or records

Organizational personnel with configuration management responsibilities

organizational personnel with information security responsibilities

system/network administrators

Organizational processes for managing baseline configurations

automated mechanisms implementing baseline configuration maintenance

### CM-2 (3): Retention of Previous Configurations

Retain the number of previous baseline configuration versions to be retained is defined; of previous versions of baseline configurations of the system to support rollback.

Retaining previous versions of baseline configurations to support rollback include hardware, software, firmware, configuration files, configuration records, and associated documentation.

 the number of previous baseline configuration versions to be retained is defined; of previous baseline configuration version(s) of the system is/are retained to support rollback.

Configuration management policy

procedures addressing the baseline configuration of the system

configuration management plan

system architecture and configuration documentation

system configuration settings and associated documentation

copies of previous baseline configuration versions

system security plan

other relevant documents or records

Organizational personnel with configuration management responsibilities

organizational personnel with information security responsibilities

system/network administrators

Organizational processes for managing baseline configurations

### CM-2 (7): Configure Systems and Components for High-risk Areas

Issue the systems or system components to be issued when individuals travel to high-risk areas are defined; with configurations for systems or system components to be issued when individuals travel to high-risk areas are defined; to individuals traveling to locations that the organization deems to be of significant risk; and

Apply the following controls to the systems or components when the individuals return from travel: the controls to be applied when the individuals return from travel are defined;.

When it is known that systems or system components will be in high-risk areas external to the organization, additional controls may be implemented to counter the increased threat in such areas. For example, organizations can take actions for notebook computers used by individuals departing on and returning from travel. Actions include determining the locations that are of concern, defining the required configurations for the components, ensuring that components are configured as intended before travel is initiated, and applying controls to the components after travel is completed. Specially configured notebook computers include computers with sanitized hard drives, limited applications, and more stringent configuration settings. Controls applied to mobile devices upon return from travel include examining the mobile device for signs of physical tampering and purging and reimaging disk drives. Protecting information that resides on mobile devices is addressed in the [MP](#mp) (Media Protection) family.

the systems or system components to be issued when individuals travel to high-risk areas are defined; with configurations for systems or system components to be issued when individuals travel to high-risk areas are defined; are issued to individuals traveling to locations that the organization deems to be of significant risk;

 the controls to be applied when the individuals return from travel are defined; are applied to the systems or system components when the individuals return from travel.

Configuration management policy

configuration management plan

procedures addressing the baseline configuration of the system

procedures addressing system component installations and upgrades

system architecture and configuration documentation

system configuration settings and associated documentation

system component inventory

records of system baseline configuration reviews and updates

system component installations/upgrades and associated records

change control records

system security plan

other relevant documents or records

Organizational personnel with configuration management responsibilities

organizational personnel with information security responsibilities

system/network administrators

Organizational processes for managing baseline configurations

### CM-3: Configuration Change Control

Determine and document the types of changes to the system that are configuration-controlled;

Review proposed configuration-controlled changes to the system and approve or disapprove such changes with explicit consideration for security and privacy impact analyses;

Document configuration change decisions associated with the system;

Implement approved configuration-controlled changes to the system;

Retain records of configuration-controlled changes to the system for the time period to retain records of configuration-controlled changes is defined;;

Monitor and review activities associated with configuration-controlled changes to the system; and

Coordinate and provide oversight for configuration change control activities through the configuration change control element responsible for coordinating and overseeing change control activities is defined; that convenes the frequency at which the configuration control element convenes is defined (if selected); and/orwhen configuration change conditions that prompt the configuration control element to convene are defined (if selected);.

Configuration change control for organizational systems involves the systematic proposal, justification, implementation, testing, review, and disposition of system changes, including system upgrades and modifications. Configuration change control includes changes to baseline configurations, configuration items of systems, operational procedures, configuration settings for system components, remediate vulnerabilities, and unscheduled or unauthorized changes. Processes for managing configuration changes to systems include Configuration Control Boards or Change Advisory Boards that review and approve proposed changes. For changes that impact privacy risk, the senior agency official for privacy updates privacy impact assessments and system of records notices. For new systems or major upgrades, organizations consider including representatives from the development organizations on the Configuration Control Boards or Change Advisory Boards. Auditing of changes includes activities before and after changes are made to systems and the auditing activities required to implement such changes. See also [SA-10](#sa-10).

the types of changes to the system that are configuration-controlled are determined and documented;

proposed configuration-controlled changes to the system are reviewed;

proposed configuration-controlled changes to the system are approved or disapproved with explicit consideration for security and privacy impact analyses;

configuration change decisions associated with the system are documented;

approved configuration-controlled changes to the system are implemented;

records of configuration-controlled changes to the system are retained for the time period to retain records of configuration-controlled changes is defined;;

activities associated with configuration-controlled changes to the system are monitored;

activities associated with configuration-controlled changes to the system are reviewed;

configuration change control activities are coordinated and overseen by the configuration change control element responsible for coordinating and overseeing change control activities is defined;;

the configuration control element convenes the frequency at which the configuration control element convenes is defined (if selected); and/orwhen configuration change conditions that prompt the configuration control element to convene are defined (if selected);.

Configuration management policy

procedures addressing system configuration change control

configuration management plan

system architecture and configuration documentation

change control records

system audit records

change control audit and review reports

agenda/minutes/documentation from configuration change control oversight meetings

system security plan

privacy plan

privacy impact assessments

system of records notices

other relevant documents or records

Organizational personnel with configuration change control responsibilities

organizational personnel with information security and privacy responsibilities

system/network administrators

members of change control board or similar

Organizational processes for configuration change control

mechanisms that implement configuration change control

### CM-3 (2): Testing, Validation, and Documentation of Changes

Test, validate, and document changes to the system before finalizing the implementation of the changes.

Changes to systems include modifications to hardware, software, or firmware components and configuration settings defined in [CM-6](#cm-6) . Organizations ensure that testing does not interfere with system operations that support organizational mission and business functions. Individuals or groups conducting tests understand security and privacy policies and procedures, system security and privacy policies and procedures, and the health, safety, and environmental risks associated with specific facilities or processes. Operational systems may need to be taken offline, or replicated to the extent feasible, before testing can be conducted. If systems must be taken offline for testing, the tests are scheduled to occur during planned system outages whenever possible. If the testing cannot be conducted on operational systems, organizations employ compensating controls.

changes to the system are tested before finalizing the implementation of the changes;

changes to the system are validated before finalizing the implementation of the changes;

changes to the system are documented before finalizing the implementation of the changes.

Configuration management policy

configuration management plan

procedures addressing system configuration change control

system design documentation

system architecture and configuration documentation

system configuration settings and associated documentation

test records

validation records

change control records

system audit records

system security plan

other relevant documents or records

Organizational personnel with configuration change control responsibilities

organizational personnel with information security responsibilities

system/network administrators

system developers

members of change control board or similar

Organizational processes for configuration change control

mechanisms supporting and/or implementing, testing, validating, and documenting system changes

### CM-3 (4): Security and Privacy Representatives

Require organization-defined security and privacy representatives to be members of the the configuration change control element of which the security and privacy representatives are to be members is defined;.

Information security and privacy representatives include system security officers, senior agency information security officers, senior agency officials for privacy, or system privacy officers. Representation by personnel with information security and privacy expertise is important because changes to system configurations can have unintended side effects, some of which may be security- or privacy-relevant. Detecting such changes early in the process can help avoid unintended, negative consequences that could ultimately affect the security and privacy posture of systems. The configuration change control element referred to in the second organization-defined parameter reflects the change control elements defined by organizations in [CM-3g](#cm-3_smt.g).

security representatives required to be members of the change control element are defined; are required to be members of the the configuration change control element of which the security and privacy representatives are to be members is defined;;

 privacy representatives required to be members of the change control element are defined; are required to be members of the the configuration change control element of which the security and privacy representatives are to be members is defined;.

Configuration management policy

procedures addressing system configuration change control

configuration management plan

system security plan

privacy plan

other relevant documents or records

Organizational personnel with configuration change control responsibilities

organizational personnel with information security and privacy responsibilities

members of change control board or similar

Organizational processes for configuration change control

### CM-4 (2): Verification of Controls

After system changes, verify that the impacted controls are implemented correctly, operating as intended, and producing the desired outcome with regard to meeting the security and privacy requirements for the system.

Implementation in this context refers to installing changed code in the operational system that may have an impact on security or privacy controls.

the impacted controls are implemented correctly with regard to meeting the security requirements for the system after system changes;

the impacted controls are implemented correctly with regard to meeting the privacy requirements for the system after system changes;

the impacted controls are operating as intended with regard to meeting the security requirements for the system after system changes;

the impacted controls are operating as intended with regard to meeting the privacy requirements for the system after system changes;

the impacted controls are producing the desired outcome with regard to meeting the security requirements for the system after system changes;

the impacted controls are producing the desired outcome with regard to meeting the privacy requirements for the system after system changes.

Configuration management policy

procedures addressing security impact analyses for changes to the system

procedures addressing privacy impact analyses for changes to the system

privacy risk assessment documentation

configuration management plan

security and privacy impact analysis documentation

privacy impact assessment

analysis tools and associated outputs

change control records

control assessment results

system audit records

system component inventory

system security plan

privacy plan

other relevant documents or records

Organizational personnel with responsibility for conducting security and privacy impact analyses

organizational personnel with information security and privacy responsibilities

system/network administrators

security and privacy assessors

Organizational processes for security and privacy impact analyses

mechanisms supporting and/or implementing security and privacy impact analyses of changes

### CM-7 (1): Periodic Review

Review the system the frequency at which to review the system to identify unnecessary and/or non-secure functions, ports, protocols, software, and/or services is defined; to identify unnecessary and/or nonsecure functions, ports, protocols, software, and services; and

Disable or remove organization-defined functions, ports, protocols, software, and services within the system deemed to be unnecessary and/or nonsecure.

Organizations review functions, ports, protocols, and services provided by systems or system components to determine the functions and services that are candidates for elimination. Such reviews are especially important during transition periods from older technologies to newer technologies (e.g., transition from IPv4 to IPv6). These technology transitions may require implementing the older and newer technologies simultaneously during the transition period and returning to minimum essential functions, ports, protocols, and services at the earliest opportunity. Organizations can either decide the relative security of the function, port, protocol, and/or service or base the security decision on the assessment of other entities. Unsecure protocols include Bluetooth, FTP, and peer-to-peer networking.

the system is reviewed the frequency at which to review the system to identify unnecessary and/or non-secure functions, ports, protocols, software, and/or services is defined; to identify unnecessary and/or non-secure functions, ports, protocols, software, and services:

functions to be disabled or removed when deemed unnecessary or non-secure are defined; deemed to be unnecessary and/or non-secure are disabled or removed;

 ports to be disabled or removed when deemed unnecessary or non-secure are defined; deemed to be unnecessary and/or non-secure are disabled or removed;

 protocols to be disabled or removed when deemed unnecessary or non-secure are defined; deemed to be unnecessary and/or non-secure are disabled or removed;

 software to be disabled or removed when deemed unnecessary or non-secure is defined; deemed to be unnecessary and/or non-secure is disabled or removed;

 services to be disabled or removed when deemed unnecessary or non-secure are defined; deemed to be unnecessary and/or non-secure are disabled or removed.

Configuration management policy

procedures addressing least functionality in the system

configuration management plan

system design documentation

system configuration settings and associated documentation

common secure configuration checklists

documented reviews of functions, ports, protocols, and/or services

change control records

system audit records

system security plan

other relevant documents or records

Organizational personnel with responsibilities for reviewing functions, ports, protocols, and services on the system

organizational personnel with information security responsibilities

system/network administrators

system developers

Organizational processes for reviewing or disabling functions, ports, protocols, and services on the system

mechanisms implementing review and disabling of functions, ports, protocols, and/or services

### CM-7 (2): Prevent Program Execution

Prevent program execution in accordance with policies, rules of behavior, and/or access agreements regarding software program usage and restrictions are defined (if selected); and/orrules authorizing the terms and conditions of software program usage.

Prevention of program execution addresses organizational policies, rules of behavior, and/or access agreements that restrict software usage and the terms and conditions imposed by the developer or manufacturer, including software licensing and copyrights. Restrictions include prohibiting auto-execute features, restricting roles allowed to approve program execution, permitting or prohibiting specific software programs, or restricting the number of program instances executed at the same time.

program execution is prevented in accordance with policies, rules of behavior, and/or access agreements regarding software program usage and restrictions are defined (if selected); and/orrules authorizing the terms and conditions of software program usage.

Configuration management policy

procedures addressing least functionality in the system

configuration management plan

system design documentation

system configuration settings and associated documentation

system component inventory

common secure configuration checklists

specifications for preventing software program execution

change control records

system audit records

system security plan

other relevant documents or records

Organizational personnel with information security responsibilities

system/network administrators

system developers

Organizational processes preventing program execution on the system

organizational processes for software program usage and restrictions

mechanisms preventing program execution on the system

mechanisms supporting and/or implementing software program usage and restrictions

### CM-7 (5): Authorized Software — Allow-by-exception

Identify software programs authorized to execute on the system are defined;;

Employ a deny-all, permit-by-exception policy to allow the execution of authorized software programs on the system; and

Review and update the list of authorized software programs frequency at which to review and update the list of authorized software programs is defined;.

Authorized software programs can be limited to specific versions or from a specific source. To facilitate a comprehensive authorized software process and increase the strength of protection for attacks that bypass application level authorized software, software programs may be decomposed into and monitored at different levels of detail. These levels include applications, application programming interfaces, application modules, scripts, system processes, system services, kernel functions, registries, drivers, and dynamic link libraries. The concept of permitting the execution of authorized software may also be applied to user actions, system ports and protocols, IP addresses/ranges, websites, and MAC addresses. Organizations consider verifying the integrity of authorized software programs using digital signatures, cryptographic checksums, or hash functions. Verification of authorized software can occur either prior to execution or at system startup. The identification of authorized URLs for websites is addressed in [CA-3(5)](#ca-3.5) and [SC-7](#sc-7).

software programs authorized to execute on the system are defined; are identified;

a deny-all, permit-by-exception policy to allow the execution of authorized software programs on the system is employed;

the list of authorized software programs is reviewed and updated frequency at which to review and update the list of authorized software programs is defined;.

Configuration management policy

procedures addressing least functionality in the system

configuration management plan

system design documentation

system configuration settings and associated documentation

list of software programs authorized to execute on the system

system component inventory

common secure configuration checklists

review and update records associated with list of authorized software programs

change control records

system audit records

system security plan

other relevant documents or records

Organizational personnel with responsibilities for identifying software authorized to execute on the system

organizational personnel with information security responsibilities

system/network administrators

Organizational process for identifying, reviewing, and updating programs authorized to execute on the system

organizational process for implementing authorized software policy

mechanisms supporting and/or implementing authorized software policy

### CM-8 (1): Updates During Installation and Removal

Update the inventory of system components as part of component installations, removals, and system updates.

Organizations can improve the accuracy, completeness, and consistency of system component inventories if the inventories are updated as part of component installations or removals or during general system updates. If inventories are not updated at these key times, there is a greater likelihood that the information will not be appropriately captured and documented. System updates include hardware, software, and firmware components.

the inventory of system components is updated as part of component installations;

the inventory of system components is updated as part of component removals;

the inventory of system components is updated as part of system updates.

Configuration management policy

procedures addressing system component inventory

configuration management plan

system security plan

system component inventory

inventory reviews and update records

change control records

component installation records

component removal records

system security plan

other relevant documents or records

Organizational personnel with component inventory updating responsibilities

organizational personnel with information security responsibilities

system/network administrators

Organizational processes for updating the system component inventory

mechanisms supporting and/or implementing system component inventory updates

### CM-8 (3): Automated Unauthorized Component Detection

Detect the presence of unauthorized hardware, software, and firmware components within the system using organization-defined automated mechanisms frequency at which automated mechanisms are used to detect the presence of unauthorized system components within the system is defined; ; and

Take the following actions when unauthorized components are detected: disable network access by unauthorized components, isolate unauthorized components, and/or notify personnel or roles to be notified when unauthorized components are detected is/are defined (if selected);.

Automated unauthorized component detection is applied in addition to the monitoring for unauthorized remote connections and mobile devices. Monitoring for unauthorized system components may be accomplished on an ongoing basis or by the periodic scanning of systems for that purpose. Automated mechanisms may also be used to prevent the connection of unauthorized components (see [CM-7(9)](#cm-7.9) ). Automated mechanisms can be implemented in systems or in separate system components. When acquiring and implementing automated mechanisms, organizations consider whether such mechanisms depend on the ability of the system component to support an agent or supplicant in order to be detected since some types of components do not have or cannot support agents (e.g., IoT devices, sensors). Isolation can be achieved , for example, by placing unauthorized system components in separate domains or subnets or quarantining such components. This type of component isolation is commonly referred to as "sandboxing." 

the presence of unauthorized hardware within the system is detected using automated mechanisms used to detect the presence of unauthorized hardware within the system are defined; frequency at which automated mechanisms are used to detect the presence of unauthorized system components within the system is defined;;

the presence of unauthorized software within the system is detected using automated mechanisms used to detect the presence of unauthorized software within the system are defined; frequency at which automated mechanisms are used to detect the presence of unauthorized system components within the system is defined;;

the presence of unauthorized firmware within the system is detected using automated mechanisms used to detect the presence of unauthorized firmware within the system are defined; frequency at which automated mechanisms are used to detect the presence of unauthorized system components within the system is defined;;

disable network access by unauthorized components, isolate unauthorized components, and/or notify personnel or roles to be notified when unauthorized components are detected is/are defined (if selected); are taken when unauthorized hardware is detected;

 disable network access by unauthorized components, isolate unauthorized components, and/or notify personnel or roles to be notified when unauthorized components are detected is/are defined (if selected); are taken when unauthorized software is detected;

 disable network access by unauthorized components, isolate unauthorized components, and/or notify personnel or roles to be notified when unauthorized components are detected is/are defined (if selected); are taken when unauthorized firmware is detected.

Configuration management policy

procedures addressing system component inventory

configuration management plan

system design documentation

system security plan

system component inventory

change control records

alerts/notifications of unauthorized components within the system

system monitoring records

system maintenance records

system audit records

system security plan

other relevant documents or records

Organizational personnel with component inventory management responsibilities

organizational personnel with responsibilities for managing the automated mechanisms implementing unauthorized system component detection

organizational personnel with information security responsibilities

system/network administrators

system developers

Organizational processes for detection of unauthorized system components

organizational processes for taking action when unauthorized system components are detected

automated mechanisms supporting and/or implementing the detection of unauthorized system components

automated mechanisms supporting and/or implementing actions taken when unauthorized system components are detected

### CM-9: Configuration Management Plan

Develop, document, and implement a configuration management plan for the system that:

Addresses roles, responsibilities, and configuration management processes and procedures;

Establishes a process for identifying configuration items throughout the system development life cycle and for managing the configuration of the configuration items;

Defines the configuration items for the system and places the configuration items under configuration management;

Is reviewed and approved by personnel or roles to review and approve the configuration management plan is/are defined; ; and

Protects the configuration management plan from unauthorized disclosure and modification.

Configuration management activities occur throughout the system development life cycle. As such, there are developmental configuration management activities (e.g., the control of code and software libraries) and operational configuration management activities (e.g., control of installed components and how the components are configured). Configuration management plans satisfy the requirements in configuration management policies while being tailored to individual systems. Configuration management plans define processes and procedures for how configuration management is used to support system development life cycle activities.

Configuration management plans are generated during the development and acquisition stage of the system development life cycle. The plans describe how to advance changes through change management processes; update configuration settings and baselines; maintain component inventories; control development, test, and operational environments; and develop, release, and update key documents.

Organizations can employ templates to help ensure the consistent and timely development and implementation of configuration management plans. Templates can represent a configuration management plan for the organization with subsets of the plan implemented on a system by system basis. Configuration management approval processes include the designation of key stakeholders responsible for reviewing and approving proposed changes to systems, and personnel who conduct security and privacy impact analyses prior to the implementation of changes to the systems. Configuration items are the system components, such as the hardware, software, firmware, and documentation to be configuration-managed. As systems continue through the system development life cycle, new configuration items may be identified, and some existing configuration items may no longer need to be under configuration control.

a configuration management plan for the system is developed and documented;

a configuration management plan for the system is implemented;

the configuration management plan addresses roles;

the configuration management plan addresses responsibilities;

the configuration management plan addresses configuration management processes and procedures;

the configuration management plan establishes a process for identifying configuration items throughout the system development life cycle;

the configuration management plan establishes a process for managing the configuration of the configuration items;

the configuration management plan defines the configuration items for the system;

the configuration management plan places the configuration items under configuration management;

the configuration management plan is reviewed and approved by personnel or roles to review and approve the configuration management plan is/are defined;;

the configuration management plan is protected from unauthorized disclosure;

the configuration management plan is protected from unauthorized modification.

Configuration management policy

procedures addressing configuration management planning

configuration management plan

system design documentation

system security plan

privacy plan

other relevant documents or records

Organizational personnel with responsibilities for developing the configuration management plan

organizational personnel with responsibilities for implementing and managing processes defined in the configuration management plan

organizational personnel with responsibilities for protecting the configuration management plan

organizational personnel with information security and privacy responsibilities

system/network administrators

Organizational processes for developing and documenting the configuration management plan

organizational processes for identifying and managing configuration items

organizational processes for protecting the configuration management plan

mechanisms implementing the configuration management plan

mechanisms for managing configuration items

mechanisms for protecting the configuration management plan

### CM-12: Information Location

Identify and document the location of information for which the location is to be identified and documented is defined; and the specific system components on which the information is processed and stored;

Identify and document the users who have access to the system and system components where the information is processed and stored; and

Document changes to the location (i.e., system or system components) where the information is processed and stored.

Information location addresses the need to understand where information is being processed and stored. Information location includes identifying where specific information types and information reside in system components and how information is being processed so that information flow can be understood and adequate protection and policy management provided for such information and system components. The security category of the information is also a factor in determining the controls necessary to protect the information and the system component where the information resides (see [FIPS 199](#628d22a1-6a11-4784-bc59-5cd9497b5445) ). The location of the information and system components is also a factor in the architecture and design of the system (see [SA-4](#sa-4), [SA-8](#sa-8), [SA-17](#sa-17)).

the location of information for which the location is to be identified and documented is defined; is identified and documented;

the specific system components on which information for which the location is to be identified and documented is defined; is processed are identified and documented;

the specific system components on which information for which the location is to be identified and documented is defined; is stored are identified and documented;

the users who have access to the system and system components where information for which the location is to be identified and documented is defined; is processed are identified and documented;

the users who have access to the system and system components where information for which the location is to be identified and documented is defined; is stored are identified and documented;

changes to the location (i.e., system or system components) where information for which the location is to be identified and documented is defined; is processed are documented;

changes to the location (i.e., system or system components) where information for which the location is to be identified and documented is defined; is stored are documented.

Configuration management policy

procedures addressing identification and documentation of information location

configuration management plan

system design documentation

system architecture documentation

PII inventory documentation

data mapping documentation

audit records

list of users with system and system component access

change control records

system component inventory

system security plan

privacy plan

other relevant documents or records

Organizational personnel with responsibilities for managing information location and user access to information

organizational personnel with responsibilities for operating, using, and/or maintaining the system

organizational personnel with information security and privacy responsibilities

system/network administrators

system developers

Organizational processes governing information location

mechanisms enforcing policies and methods for governing information location

### CM-12 (1): Automated Tools to Support Information Location

Use automated tools to identify information to be protected is defined by information type; on system components where the information is located are defined; to ensure controls are in place to protect organizational information and individual privacy.

The use of automated tools helps to increase the effectiveness and efficiency of the information location capability implemented within the system. Automation also helps organizations manage the data produced during information location activities and share such information across the organization. The output of automated information location tools can be used to guide and inform system architecture and design decisions.

automated tools are used to identify information to be protected is defined by information type; on system components where the information is located are defined; to ensure that controls are in place to protect organizational information and individual privacy.

Configuration management policy

procedures addressing identification and documentation of information location

configuration management plan

system design documentation

PII inventory documentation

data mapping documentation

change control records

system component inventory

system security plan

privacy plan

other relevant documents or records

Organizational personnel with responsibilities for managing information location

organizational personnel with information security responsibilities

system/network administrators

system developers

Organizational processes governing information location

automated mechanisms enforcing policies and methods for governing information location

automated tools used to identify information on system components

