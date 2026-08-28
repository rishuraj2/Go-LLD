# Functional Requirements
- Support for "Email", "SMS" and "Push" notification types.
- Retry sending failed notifications with maximum number of retries and delay between attempts.
- Sending notification should be asynchronous.
- Support sending only one notification per request.
- Notification may contain a subject (optional) and a message body (mandatory).

# Non-functional Requirements
- The system should follow object-oriented design with clear seperation of concerns.
- Design should be extensible allowing future support for new notification types (e.g., whatsApp, Slack).
- Design should be asynchronous, using a thread pool to manage parallel sending.

# Core Entities
## Enums
- Notification_Type

## Data Class
- Notification_Config
- Recipient

## Core Class
- Job_Queue
- Retry_Queue
- Job_Pool
- Notification

# Facade
- NotificationService

# Responsibilities
- **NotificationType**: Enum having "**EMAIL**", "**SMS**" and "**PUSH**" and value
- **Notification_Config**: It has details about **MAX_RETRIES** and **TIME_INTERVAL** in between each try.
- **Recipient**: It is a unique identifier that determines whom to send the **Notification**.
- **JobQueue**: It is a **FIFO queue** where all the **Notifications** are first sent and **JobPool** is supposed to pick task for it.
- **RetryQueue**: It is a **Priority queue** where Notifications are sent if they fail in the first go. Here, jobs will be ordered on the basis of one that needs to be retried first. Check out min-heap concept.
- **JobPool**: It consists of worker threads waiting to pick the task from either **JobQueue** or **RetryQueue**.
- **Notification**: It holds the data for notification.