import { PaymentMethodsInterface } from "./IPaymentmethod";
import { DeliveryTypesInterface } from "./IDeliverytype";
import { OrdersInterface } from "./IOrder";

export interface PaymentInterface {
  ID: string,
  Phone: string,
  PaymentTime: Date,
  DeliveryTypeID: number,
  DeliveryType: DeliveryTypesInterface,
  PaymentMethodID: number,
  PaymentMethod: PaymentMethodsInterface,
  OrderID: number,
  Order: OrdersInterface,
}