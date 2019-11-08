/*
 * @Author: berryberry
 * @since: 2019-11-08 16:59:10
 * @LastModifiedBy: berryberry
 * @LastModifiedTime: Do not edit
 */
package mongorm

type NamableCollection interface {
	CollectionName() string
}
